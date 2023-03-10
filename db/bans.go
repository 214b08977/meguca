package db

import (
	"database/sql"
	"sync"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/bakape/meguca/auth"
	"github.com/bakape/meguca/common"
	"github.com/go-playground/log"
)

var (
	// board: IP: IsBanned
	banCache = map[string]map[string]bool{}
	bansMu   sync.RWMutex
	banTypes = map[common.ModerationAction]string{
		common.BanPost:       "classic",
		common.ShadowBinPost: "shadow",
	}
)

func writeBan(tx *sql.Tx, ip string, entry auth.ModLogEntry) (err error) {
	_, err = sq.Insert("bans").
		Columns("ip", "board", "forPost", "reason", "by", "type", "expires").
		Values(ip, entry.Board, entry.ID, entry.Data, entry.By,
			banTypes[entry.Type],
			time.Now().UTC().Add(time.Second*time.Duration(entry.Length))).
		RunWith(tx).
		Exec()
	if err != nil {
		return
	}

	return logModeration(tx, entry)
}

// Propagate ban updates through DB and disconnect all banned IPs
func propagateBans(tx *sql.Tx, board string, ip string) (err error) {
	_, err = tx.Exec(`notify bans_updated`)
	if err != nil {
		return
	}
	if !common.IsTest {
		auth.DisconnectByBoardAndIP(ip, board)
	}
	return
}

// Automatically bans an IP
func SystemBan(ip, board, reason string, length time.Duration) (err error) {
	return InTransaction(false, func(tx *sql.Tx) error {
		return systemBanTx(tx, ip, reason, length)
	})
}

func systemBanTx(tx *sql.Tx, ip, reason string, length time.Duration,
) (
	err error,
) {
	err = writeBan(tx, ip, auth.ModLogEntry{
		ModerationEntry: common.ModerationEntry{
			Type:   common.BanPost,
			Data:   reason,
			By:     "system",
			Length: uint64(length / time.Second),
		},
		Board: "all",
	})
	if err != nil {
		return
	}
	err = propagateBans(tx, "all", ip)
	return
}

// Ban IP from accessing a specific board. Need to target a post.
func Ban(
	tx *sql.Tx, board, reason, by string, length time.Duration,
	id uint64, banType common.ModerationAction,
) (
	err error,
) {
	ip, err := GetIP(id)
	switch err {
	case nil:
		if ip == "" { // Post already cleared of IP
			return
		}
	case sql.ErrNoRows:
		return nil
	default:
		return
	}

	// Write ban messages to posts and ban table
	writeBan(tx, ip, auth.ModLogEntry{
		ModerationEntry: common.ModerationEntry{
			Type:   banType,
			Length: uint64(length / time.Second),
			By:     by,
			Data:   reason,
		},
		Board: board,
		ID:    id,
	})
	if err != nil {
		return
	}

	if banType == common.BanPost {
		return propagateBans(tx, board, ip)
	}
	return
}

// Unban lifts a ban from a specific post on a specific board
func Unban(board string, id uint64, by string) error {
	return InTransaction(false, func(tx *sql.Tx) (err error) {
		_, err = sq.Delete("bans").
			Where("board = ? and forPost = ?", board, id).
			RunWith(tx).
			Exec()
		if err != nil {
			return
		}
		exists := false
		err = sq.Select("true").
			From("posts").
			Where("id = ? and board = ?", id, board).
			QueryRow().
			Scan(&exists)
		if err != nil && err != sql.ErrNoRows {
			return
		}
		if exists {
			err = logModeration(tx, auth.ModLogEntry{
				ModerationEntry: common.ModerationEntry{
					Type: common.UnbanPost,
					By:   by,
				},
				Board: board,
				ID:    id,
			})
			if err != nil {
				return
			}
		}
		_, err = tx.Exec("notify bans_updated")
		return
	})
}

func loadBans() error {
	if err := RefreshBanCache(); err != nil {
		return err
	}
	return Listen("bans_updated", func(_ string) error {
		return RefreshBanCache()
	})
}

func selectBans(colums ...string) squirrel.SelectBuilder {
	return sq.Select(colums...).
		Options("distinct on (ip, board)").
		From("bans").
		Where("expires > now() at time zone 'utc' and type = 'classic'").
		OrderBy("ip", "board", "expires desc")
}

// RefreshBanCache loads up to date bans from the database and caches them in
// memory
func RefreshBanCache() (err error) {
	bans := make([]auth.Ban, 0, 16)
	err = queryAll(selectBans("ip", "board"), func(r *sql.Rows) error {
		var b auth.Ban
		err := r.Scan(&b.IP, &b.Board)
		if err != nil {
			return err
		}
		bans = append(bans, b)
		return nil
	})
	if err != nil {
		return
	}

	new := map[string]map[string]bool{}
	for _, b := range bans {
		board, ok := new[b.Board]
		if !ok {
			board = map[string]bool{}
			new[b.Board] = board
		}
		board[b.IP] = true
	}

	bansMu.Lock()
	banCache = new
	bansMu.Unlock()

	return
}

// IsBanned checks, if the IP is banned on the target board or globally.
// Returns ErrBanned error if banned and boolean for if the ban is global
func IsBanned(board, ip string) (bool, error) {
	bansMu.RLock()
	defer bansMu.RUnlock()
	global := banCache["all"]
	ips := banCache[board]
	isGlobal := false

	if (global != nil && global[ip]) || (ips != nil && ips[ip]) {
		// Need to assert ban has not expired and cache is invalid

		r, err := selectBans("board").Where("ip = ?", ip).Query()
		if err != nil {
			return isGlobal, err
		}
		defer r.Close()

		var (
			resBoard string
			matched  = false
		)
		for r.Next() {
			err = r.Scan(&resBoard)
			if err != nil {
				return isGlobal, err
			}
			if resBoard == "all" {
				matched = true
				isGlobal = true
				break
			} else if resBoard == board {
				matched = true
				break
			}
		}
		err = r.Err()
		if err != nil {
			return isGlobal, err
		}

		if matched {
			// Also refresh the cache to keep stale positives from
			// triggering a check again
			if !common.IsTest {
				go func() {
					err := RefreshBanCache()
					if err != nil {
						log.Error(err)
					}
				}()
			}

			return isGlobal, common.ErrBanned
		}
		return isGlobal, nil
	}

	return isGlobal, nil
}

// GetBanInfo retrieves information about a specific ban
func GetBanInfo(ip, board string) (b auth.BanRecord, err error) {
	err = sq.Select("ip", "board", "forPost", "reason", "by", "expires").
		From("bans").
		Where(
			`expires >= now() at time zone 'utc'
					and ip = ?
					and board = ?
					and type = 'classic'`,
			ip, board).
		QueryRow().
		Scan(&b.IP, &b.Board, &b.ForPost, &b.Reason, &b.By, &b.Expires)
	return
}

// GetBoardBans gets all bans on a specific board. "all" counts as a valid board value.
func GetBoardBans(board string) (b []auth.BanRecord, err error) {
	b = make([]auth.BanRecord, 0, 64)
	rec := auth.BanRecord{
		Ban: auth.Ban{
			Board: board,
		},
	}
	err = queryAll(
		sq.Select("ip", "forPost", "reason", "by", "expires", "type").
			From("bans").
			Where("expires >= now() at time zone 'utc' and board = ?", board),
		func(r *sql.Rows) (err error) {
			err = r.Scan(&rec.IP, &rec.ForPost, &rec.Reason, &rec.By,
				&rec.Expires, &rec.Type)
			if err != nil {
				return
			}
			b = append(b, rec)
			return
		},
	)
	return
}

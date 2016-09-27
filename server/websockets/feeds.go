// Thread update feed managment

package websockets

import (
	"log"
	"sync"

	"github.com/bakape/meguca/db"
	r "github.com/dancannon/gorethink"
)

var (
	// Precompiled query for extracting only the changed fields from the
	// replication log feed
	formatUpdateFeed = r.Branch(
		r.Row.HasFields("old_val"),
		r.Row.
			Field("new_val").
			Field("log").
			Slice(r.Row.Field("old_val").Field("log").Count()).
			Default(nil), // Thread deleted
		r.Row.Field("new_val").Field("log").Count(), // Initial counter
	)

	// Contains and manages all active update feeds
	feeds = feedContainer{
		// 100 len map to avoid some realocation as the server starts
		feeds: make(map[int64]*updateFeed, 100),
	}
)

// Container for holding and managing client<->update-feed interaction
type feedContainer struct {
	sync.RWMutex
	feeds map[int64]*updateFeed
}

// A feed with syncronisation logic of a certain thread
type updateFeed struct {
	id      int64         // Thread ID
	ctr     uint64        // Feed progress counter
	clients []*Client     // Subscribed clients
	Add     chan *Client  // Add a client to u
	Remove  chan *Client  // Remove a client from u
	close   chan struct{} // Close database change feed
	read    chan [][]byte // Read from database change feed
}

// Add a client to an existing update feed or create a new one, if it does not
// exist yet
func (f *feedContainer) Add(id int64, cl *Client) (*updateFeed, error) {
	f.Lock()
	defer f.Unlock()

	feed := f.feeds[id]
	if feed != nil {
		feed.Add <- cl
		return feed, nil
	}

	var err error
	feed, err = newUpdateFeed(id)
	if err != nil {
		close(feed.close)
		return nil, err
	}

	f.feeds[id] = feed
	feed.Add <- cl

	return feed, nil
}

// Remove an updateFeed from f. Should only be called by the feed itself.
func (f *feedContainer) Remove(id int64) {
	f.Lock()
	defer f.Unlock()
	delete(f.feeds, id)
}

// Stop and remove all existing feeds. Used only in tests.
func (f *feedContainer) Clear() {
	f.Lock()
	defer f.Unlock()
	for id, feed := range f.feeds {
		select {
		case <-feed.close:
		default:
			close(feed.close)
		}
		delete(f.feeds, id)
	}
}

// Create a new updateFeed and sync it to the database
func newUpdateFeed(id int64) (*updateFeed, error) {
	feed := updateFeed{
		id:      id,
		clients: make([]*Client, 0, 1),
		close:   make(chan struct{}),
		read:    make(chan [][]byte),
		Add:     make(chan *Client),
		Remove:  make(chan *Client),
	}

	cursor, err := feed.streamUpdates()
	if err != nil {
		return nil, err
	}

	go feed.Listen(cursor)

	return &feed, nil
}

// Start listening for updates from database and client requests
func (u *updateFeed) Listen(cursor *r.Cursor) {

	defer func() {
		err := cursor.Err()
		if err == nil {
			err = cursor.Close()
		}
		if err != nil {
			log.Printf("update feed: %s\n", err)
		}

		feeds.Remove(u.id)
	}()

	for {
		select {

		// Add client
		case client := <-u.Add:
			u.clients = append(u.clients, client)
			err := client.sendMessage(messageSynchronise, u.ctr)
			if err != nil {
				client.Close(err)
			}

		// Remove client or close feed, if no clients would remain
		case client := <-u.Remove:
			if len(u.clients) == 1 {
				u.clients = nil
				return
			}
			for i, cl := range u.clients {
				if cl == client {
					copy(u.clients[i:], u.clients[i+1:])
					u.clients[len(u.clients)-1] = nil
					u.clients = u.clients[:len(u.clients)-1]
					break
				}
			}

		// Forward updates to clients
		case updates := <-u.read:
			if updates == nil { // Thread deleted
				return
			}
			concat := ConcatMessages(updates)
			for _, client := range u.clients {
				if err := client.send(concat); err != nil {
					client.Close(err)
				}
			}

		// Feed terminated externally
		case <-u.close:
			return
		}
	}
}

// ConcatMessages concatenate multiple feed messages into a single one to reduce
// transport overhead
func ConcatMessages(msgs [][]byte) []byte {
	if len(msgs) == 1 {
		return msgs[0]
	}

	// Calculate capacity
	cap := 2 + len(msgs)
	for _, msg := range msgs {
		cap += len(msg)
	}

	buf := make([]byte, 2, cap)
	buf[0] = 52
	buf[1] = 50
	for i, msg := range msgs {
		if i != 0 {
			buf = append(buf, '\u0000') // Delimit with null bytes
		}
		buf = append(buf, msg...)
	}

	return buf
}

// StreamUpdates produces a stream of the replication log updates for the
// specified thread and sends it on read. Close the cursor to stop receiving
// updates. The intial contents of the log are assigned emediately.
func (u *updateFeed) streamUpdates() (*r.Cursor, error) {
	cursor, err := r.
		Table("threads").
		Get(u.id).
		Changes(r.ChangesOpts{
			IncludeInitial: true,
			Squash:         0.2, // Perform at most every 0.2 seconds
		}).
		Map(formatUpdateFeed).
		Run(db.RSession)
	if err != nil {
		return nil, err
	}

	if !cursor.Next(&u.ctr) {
		if err := cursor.Err(); err != nil {
			return nil, err
		}
	}

	cursor.Listen(u.read)

	return cursor, nil
}

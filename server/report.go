package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/bakape/meguca/auth"
	"github.com/bakape/meguca/common"
	"github.com/bakape/meguca/db"
	"github.com/bakape/meguca/templates"
)

type reportEvent struct {
	Post   uint64
	Reason string
}

// Report a post for rule violations
func report(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, jsonLimit)
	err := r.ParseMultipartForm(0)
	if err != nil {
		httpError(w, r, common.StatusError{err, 400})
		return
	}
	f := r.Form

	ip, err := auth.GetIP(r)
	if err != nil {
		httpError(w, r, common.StatusError{err, 400})
		return
	}
	var session auth.Base64Token
	err = session.EnsureCookie(w, r)
	if err != nil {
		httpError(w, r, common.StatusError{err, 400})
		return
	}

	has, err := db.SolvedCaptchaRecently(session, time.Minute)
	if err != nil {
		httpError(w, r, err)
		return
	}
	if !has {
		httpError(w, r, errInvalidCaptcha)
		return
	}

	target, err := strconv.ParseUint(f.Get("target"), 10, 64)
	if err != nil {
		httpError(w, r, common.StatusError{err, 400})
		return
	}

	board, err := db.GetPostBoard(target)
	if err != nil {
		httpError(w, r, err)
		return
	}
	if !auth.IsNonMetaBoard(board) {
		httpError(w, r, errInvalidBoardName)
		return
	}
	if !assertNotBanned(w, r, board) {
		return
	}

	reason := f.Get("reason")
	if len(reason) > common.MaxLenReason {
		httpError(w, r, errReasonTooLong)
		return
	}

	illegal := (f.Get("illegal") == "on")
	err = db.Report(target, board, reason, ip, illegal)
	if err != nil {
		httpError(w, r, err)
		return
	}

	data, err := json.Marshal(reportEvent{Post: target, Reason: reason})
	if err != nil {
		httpError(w, r, err)
		return
	}

	var dst string
	if illegal {
		dst = "all"
	} else {
		dst = board
	}
	SSEBroker.Event <- ServerEvent{
		Destination: "/html/reports/" + dst,
		Data:        data,
	}
}

// Render post reporting form
func reportForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(extractParam(r, "id"), 10, 64)
	if err != nil {
		httpError(w, r, common.StatusError{err, 400})
		return
	}
	setHTMLHeaders(w)
	templates.WriteReportForm(w, id)
}

// Render a list of reports for the board
func reportList(w http.ResponseWriter, r *http.Request) {
	board := extractParam(r, "board")
	if !auth.IsBoard(board) {
		text404(w)
		return
	}
	rep, err := db.GetReports(board)
	if err != nil {
		httpError(w, r, err)
		return
	}
	setHTMLHeaders(w)
	templates.WriteReportList(w, rep)
}

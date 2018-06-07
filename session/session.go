package session

import (
	"time"
	"crypto/sha1"
	"encoding/hex"
)

type session struct {
	id		string
	username 	string
	isManager	bool
	expire		time.Time
}

func newSessionId(username string) string {
	h := sha1.New()
	h.Write([]byte(username + time.Now().String()))
	return hex.EncodeToString(h.Sum(nil))
}

func newSession(username string, isManager bool) *session {
	id := newSessionId(username)
	return &session {
		id: id,
		username: username,
		isManager: isManager,
		expire: time.Now().Add(time.Hour),
	}
}

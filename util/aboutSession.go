package util

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("01234567890123456789012345678901")

	Store = sessions.NewCookieStore(key)
)

func makeSession(r *http.Request) *sessions.Session {
	session, _ := Store.Get(nil, "session-name")
	return session
}

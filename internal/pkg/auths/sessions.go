package auths

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	SESSION_KEY  = "ninja-go"
	SESSION_NAME = "ninja-session"
)

var (
	store = sessions.NewCookieStore([]byte(SESSION_KEY))
)

func GetSession(r *http.Request, name string) (*sessions.Session, error) {
	return store.Get(r, name)
}

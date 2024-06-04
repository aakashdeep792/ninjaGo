package auths

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	SESSION_KEY  = "ninja-go"
	SESSION_NAME = "ninja-session"
	SESSION_VALUE_KEY = "user"
)

var (
	cookieStore = sessions.NewCookieStore([]byte(SESSION_KEY))
)

func GetSession(r *http.Request, name string) (*sessions.Session, error) {
	return cookieStore.Get(r, name)
}

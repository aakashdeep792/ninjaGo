package handlers

import (
	"fmt"
	"net/http"
	"ninjaGo/internal/pkg/auths"

	"github.com/gorilla/sessions"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login-handler")
	session, err := auths.GetSession(r, auths.SESSION_NAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session.Options = &sessions.Options{
		Path: "/",
		// MaxAge:   86400 * 7,
		MaxAge:   5,
		HttpOnly: true,
	}

	session.Values["user"] = "aakash-signup"

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

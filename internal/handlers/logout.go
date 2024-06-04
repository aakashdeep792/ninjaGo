package handlers

import (
	"fmt"
	"net/http"
	"ninjaGo/internal/logger"
	"ninjaGo/internal/pkg/auths"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := auths.GetSession(r, auths.SESSION_NAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Debugf("session value for key %v is %v", auths.SESSION_VALUE_KEY, session.Values[auths.SESSION_VALUE_KEY])
	// Expire the session by setting the expiration time to a past time
	session.Options.MaxAge = -1 // Set to a past time to expire the session

	if err = session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Logout successful. Session expired.")
}

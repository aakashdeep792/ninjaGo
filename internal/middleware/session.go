package middleware

import (
	"fmt"
	"net/http"
	"ninjaGo/internal/env"
	"ninjaGo/internal/logger"
	"ninjaGo/internal/pkg/auths"
	"ninjaGo/internal/utils"

	"github.com/gorilla/sessions"
)

func SessionMiddleware(next http.Handler) http.Handler {
	skipURLS := env.GetSkipURLList()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if utils.SkipURL(r.URL.RequestURI(), skipURLS) {
			logger.Debug("check skipped for route", r.URL.RequestURI())
			next.ServeHTTP(w, r)
			return
		}

		session, err := auths.GetSession(r, auths.SESSION_NAME)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		logger.Debug("session value key :", session.Values[auths.SESSION_VALUE_KEY])
		sessionVal := session.Values[auths.SESSION_VALUE_KEY]
		// check if session exist and is valid
		if session.IsNew || sessionVal == nil {
			logger.Log("request unauthorized")
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		session.Options = &sessions.Options{
			Path:   "/",
			MaxAge: 86400 * 7,
			// MaxAge:   5,
			HttpOnly: true,
		}

		session.Values[auths.SESSION_VALUE_KEY] = sessionVal
		logger.Debug("session value key :", session.Values[auths.SESSION_VALUE_KEY])

		// save the session data
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		next.ServeHTTP(w, r)
		fmt.Println("2", session)
	})
}

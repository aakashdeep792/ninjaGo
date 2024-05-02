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
		sessionValueKey := "user"

		fmt.Println("middleware")

		if utils.SkipURL(r.URL.RequestURI(), skipURLS) {
			logger.Debug("skip-middleware", r.URL.RequestURI())
			next.ServeHTTP(w, r)
			return
		}

		session, err := auths.GetSession(r, auths.SESSION_KEY)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		// check if session exist and is valid
		if session.IsNew || session.Values[sessionValueKey] == nil {
			fmt.Println("unauthorized")
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		session.Options = &sessions.Options{
			Path: "/",
			// MaxAge:   86400 * 7,
			MaxAge:   5,
			HttpOnly: true,
		}

		session.Values[sessionValueKey] = "aakash@user"
		fmt.Println(session)

		// save the session data
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		next.ServeHTTP(w, r)
		fmt.Println("2", session)
	})
}

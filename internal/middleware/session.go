package middleware

import (
	"fmt"
	"net/http"
	"ninjaGo/internal/pkg/auths"
	"strings"

	"github.com/gorilla/sessions"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		if strings.Contains(r.URL.RequestURI(), "login") {
			fmt.Println("skip-middleware", r.URL.RequestURI())
			next.ServeHTTP(w, r)
			return
		}

		session, err := auths.GetSession(r, auths.SESSION_KEY)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		session.Options = &sessions.Options{
			Path: "/",
			// MaxAge:   86400 * 7,
			MaxAge:   5,
			HttpOnly: true,
		}

		session.Values["user"] = "aakash@user"
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

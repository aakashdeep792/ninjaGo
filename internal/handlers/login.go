package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ninjaGo/internal/logger"
	"ninjaGo/internal/models"
	"ninjaGo/internal/pkg/auths"

	"github.com/gorilla/sessions"
)

// type requestData struct {
// 	User     string `json:"user"`
// 	Password string `json:"password"`
// }

const (
	sessionValueKey = "user"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := auths.GetSession(r, auths.SESSION_NAME)
	if err != nil {
		logger.Logf("Error in getting session: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if session exist and is valid
	if !(session.IsNew || session.Values[sessionValueKey] == nil) {
		// redirect
		tmp := session.Values[sessionValueKey]
		fmt.Fprintf(w, "You are already logged in %s!", tmp)
		return
	}

	var requestData models.Credential
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := auths.ValidateCredential(requestData.User, requestData.Password); err != nil {
		logger.Logf("Authentication failed: %v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
		// MaxAge:   5,
		HttpOnly: true,
	}
	// update user value in session
	session.Values[auths.SESSION_VALUE_KEY] = requestData.User
	logger.Debugf("session value for key %v is %v", auths.SESSION_VALUE_KEY, session.Values[auths.SESSION_VALUE_KEY])

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Login successful, Welcome,%s!", requestData.User)))
	w.WriteHeader(http.StatusOK)
}

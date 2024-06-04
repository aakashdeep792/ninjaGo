package handlers

import (
	"encoding/json"
	"net/http"
	"ninjaGo/internal/models"
	"ninjaGo/internal/store"
)

type UserInfoWithPassword struct {
	// User     string `json:"user"`
	// Password string `json:"password"`
	// ConfirmPass string `json:"confirmPassword"`

	models.Credential
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var info UserInfoWithPassword
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.AddUserPass(info.Credential)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//  Update the statuscode then add the message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Sign Up succesful"))
}

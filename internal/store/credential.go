package store

import "ninjaGo/internal/models"

var cred = map[string]string{"aakash": "deep", "king": "Ram"}

func GetUserPass(usr string) *string {
	if v, ok := cred[usr]; ok {
		return &v
	}

	// user does not exist
	return nil
}

func AddUserPass(c models.Credential) error {
	cred[c.User] = c.Password
	return nil
}

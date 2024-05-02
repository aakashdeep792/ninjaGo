package auths

import (
	"errors"
	"ninjaGo/internal/store"
)

func validateCredential(usr, pass string) error {
	val := store.GetUserPass(usr)

	if val == nil {
		return errors.New("User does not exist")
	}

	if *val != pass {
		return errors.New("Incorrect password")
	}

	return nil
}

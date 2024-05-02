package store

var cred = map[string]string{"aakash": "deep", "king": "Ram"}

func GetUserPass(usr string) *string {
	if v, ok := cred[usr]; ok {
		return &v
	}
	
	// user does not exist
	return nil
}

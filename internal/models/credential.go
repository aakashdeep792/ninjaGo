package models

type Secret string
type SecretHash string

type Credential struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// func (c Secret) GetHash(pass string) SecretHash {
// 	h := fnv.New64a()
// 	h.Write([]byte(pass))
// 	return SecretHash(fmt.Sprint(h.Sum64()))
// }

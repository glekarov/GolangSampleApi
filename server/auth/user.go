package auth

import (
	"encoding/base64"
)

type User struct {
	username string
	password string
	token    string
}

func NewUser(username string, password string) *User {
	u := new(User)
	u.username = username
	u.password = password
	u.token = base64.StdEncoding.EncodeToString([]byte(username + "|" + password))
	return u
}
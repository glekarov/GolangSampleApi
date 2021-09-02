package auth

import (
	"errors"
)

const (
	invalidUserAndPasswd = "Invalid username or password"
	userAlreadyExists = "User already exists"
)

// A list of the created usres
type usersList []User

// Singleton representing the users list
var instance usersList

// valide whether the user exists or not
// returns true if the usre exist else false
func isUserExists(username string) bool {
	for _, l := range instance {
		if l.username == username {
			return true
		}
	}

	return false
}

// Get a user object by specified username
// returns User object else nil
func getUserByUsername(username string) *User {
	for _, l := range instance {
		if l.username == username {
			return &l
		}
	}

	return nil
}

// Add user to the list
// returns error if the user already exists
func AddUser(username string, password string) error {
	if isUserExists(username) {
		return errors.New(userAlreadyExists)
	}

	user := NewUser(username, password)
	instance = append(instance, *user)

	return nil
}

// Validates whether there is a user with the passed credentials
// returns error message if credentials are no valid, else nul
func ValidateUserAndPassword(username string, password string) error {
	user := getUserByUsername(username)
	if user != nil && user.password == password {
		return nil
	}

	return errors.New(invalidUserAndPasswd)
}

// Validates whether there is a user with the passed credentials
// returns error message if credentials are no valid, else nul
func ValidateUserAndToken(username string, token string) error {
	user := getUserByUsername(username)
	if user != nil && user.token == token {
		return nil
	}

	return errors.New(invalidUserAndPasswd)
}

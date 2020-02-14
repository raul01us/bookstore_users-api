package services

import (
	"github.com/raul01us/bookstore_users-api/domains/users"
	"github.com/raul01us/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// Validate the user
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// Attempt to save the user in the database
	if err := user.Save(); err != nil {
		return nil, err
	}
	// Return the new user
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	// Result is a pointer to a instance of the user
	result := &users.User{ID: userID}
	// Attempt to get the user
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

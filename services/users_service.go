package services

import (
	"github.com/raul01us/bookstore_users-api/domains/users"
	"github.com/raul01us/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

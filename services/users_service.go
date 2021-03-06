package services

import (
	"github.com/raul01us/bookstore_users-api/domains/users"
	"github.com/raul01us/bookstore_users-api/utils/crypto_utils"
	"github.com/raul01us/bookstore_users-api/utils/date_utils"
	"github.com/raul01us/bookstore_users-api/utils/errors"
)

var (
	// Create a new instance of the usersService struct
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

// Provide an interface to CRUD will be useful for mocks
type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// Validate the user
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Add the default values
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	// Attempt to save the user in the database
	if err := user.Save(); err != nil {
		return nil, err
	}
	// Return the new user
	return &user, nil
}

func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	// Result is a pointer to a instance of the user
	result := &users.User{ID: userID}
	// Attempt to get the user
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	// Attempt to get the user from DB before performing the update
	current, err := UsersService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		// Partial request update only passed values
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		// Update every field on the current user with the passed values
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	// Attempt to execute the update
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

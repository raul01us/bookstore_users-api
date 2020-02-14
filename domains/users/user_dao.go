package users

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/raul01us/bookstore_users-api/utils/date_utils"

	"github.com/raul01us/bookstore_users-api/datasources/mysql/users_db"

	"github.com/raul01us/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	// Create, validate and close the prepared statement
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// Attempt getting the user from the DB
	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if err == sql.ErrNoRows {
			return errors.NewNotFoundError("record not found")
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d", user.ID))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	// Create, validate and close the prepared statement
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	// Attempt saving the user in the database by executing the statement
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError("email already exists")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when saving the user: %s", err.Error()))
	}

	// Get the newly created user ID
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when saving the user: %s", err.Error()))
	}
	user.ID = userID
	return nil
}

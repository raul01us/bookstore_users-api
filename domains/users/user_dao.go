package users

import (
	"github.com/raul01us/bookstore_users-api/utils/mysql_utils"

	"github.com/raul01us/bookstore_users-api/utils/date_utils"

	"github.com/raul01us/bookstore_users-api/datasources/mysql/users_db"

	"github.com/raul01us/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
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
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
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
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	// Get the newly created user ID
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}
	user.ID = userID
	return nil
}

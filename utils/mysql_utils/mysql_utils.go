package mysql_utils

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/raul01us/bookstore_users-api/utils/errors"
)

func ParseError(err error) *errors.RestErr {
	// Attempt to convert the error to a MySQL error
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		// NoRows is not actually a MYSQLError
		if err == sql.ErrNoRows {
			return errors.NewNotFoundError("record not found")
		}
		// Conversion failed the error is not a MySQL one
		return errors.NewInternalServerError("error parsing database response")
	}

	// Define general messages based on MySQL error numbers
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}

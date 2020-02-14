package users

import (
	"net/http"
	"strconv"

	"github.com/raul01us/bookstore_users-api/utils/errors"

	"github.com/raul01us/bookstore_users-api/services"

	"github.com/raul01us/bookstore_users-api/domains/users"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Initialize a user
	var user users.User
	// Attempt to bind the request payload to the initialised user
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// Send the user to the service in order to be created and persisted (stored in DB)
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	// Return success
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	// Validate the user ID passed as a parameter
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	// Send the ID to the service in order to get the user
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/raul01us/bookstore_users-api/logger"
)

var router = gin.Default()

func StartApplication() {
	mapURLs()
	logger.Info("about to start the application...")
	router.Run(":8080")
}

package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	router.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "post request"})
	})

	router.Run(":8000")
}

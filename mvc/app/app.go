package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

// Temporary
func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrls()
	if err := router.Run(":8080"); err !=nil {
		panic(err)
	}
}
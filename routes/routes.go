package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Init foutes function
func Init() {
	router := gin.Default()

	router.GET("/", userReq)
	router.GET("/count", numberOfReq)

	router.NoRoute(func(c *gin.Context) {
		c.XML(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.Run(":8080")
}


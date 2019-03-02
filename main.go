package main

import (
	"github.com/gin-gonic/gin"
)


func main(){

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
		return
	})
	router.Run("127.0.0.1:6066")
}


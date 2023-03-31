package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize Router (r) with gin default settings
	r := gin.Default()
	// create a Route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // -set port to gin app and run
}

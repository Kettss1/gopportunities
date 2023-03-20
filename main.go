package main

import "github.com/gin-gonic/gin"

func main() {
	// initialize Router with gin default settings
	r := gin.Default()
	// create a Route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

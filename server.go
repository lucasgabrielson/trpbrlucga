package main

import "github.com/gin-gonic/gin"

func main() {

	// Creates a gin router with default middleware
	router := gin.Default()

	// A handler for GET request on /example
	router.GET("/hallo", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "hallo",
		}) // gin.H is a shortcut for map[string]interface{}
	})
	router.Run() // listen and serve on port 8080
}

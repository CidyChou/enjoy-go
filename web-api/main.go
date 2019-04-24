package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Defualt()

	router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })
    router.Run(":8000")
}

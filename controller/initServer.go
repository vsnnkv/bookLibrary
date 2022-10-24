package controller

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	router.GET("/book/:id")
	router.GET("/books")
	router.POST("/book")
	router.PUT("/book/:id")
	router.DELETE("/book/:id")

	router.Run(":8080")
}

package controllers

import (
	"bookLibrary/repository"
	"bookLibrary/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	bookController *BookController
}

func NewHandler(bookController *BookController) *Handler {
	return &Handler{bookController: bookController}
}
func (h *Handler) StartServer() {
	router := gin.Default()

	router.GET("/book/:id")
	router.GET("/books")
	router.POST("/book", h.bookController.AddBook)
	router.PUT("/book/:id")
	router.DELETE("/book/:id")

	router.Run(":8080")
}

func InitControllers() {
	bookStorage := repository.BookStorage{}

	bookService := services.NewBookService(&bookStorage)

	bookController := NewBookController(bookService)

	handler := NewHandler(bookController)
	handler.StartServer()
}

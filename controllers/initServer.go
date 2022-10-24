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

	router.GET("/book/:id", h.bookController.GetBook)
	router.GET("/books", h.bookController.GetBooks)
	router.POST("/book", h.bookController.AddBook)
	router.POST("/book/:id", h.bookController.UpdateBook)
	router.DELETE("/book/:id", h.bookController.DeleteBook)

	router.Run(":8080")
}

func InitControllers() {
	bookStorage := repository.BookStorage{}

	bookService := services.NewBookService(&bookStorage)

	bookController := NewBookController(bookService)

	handler := NewHandler(bookController)
	handler.StartServer()
}

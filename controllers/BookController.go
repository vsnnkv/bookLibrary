package controllers

import (
	"bookLibrary/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	bookService services.BookServiceInterface
}

func NewBookController(bookService services.BookServiceInterface) *BookController {
	return &BookController{bookService: bookService}
}

func (controller *BookController) AddBook(c *gin.Context) {
	var passedParam passedBook

	if err := c.BindJSON(&passedParam); err != nil {
		return
	}
	bookId, err := controller.bookService.AddBook(passedParam.Name, passedParam.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"bookId": bookId})
	}
}

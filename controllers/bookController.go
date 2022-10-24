package controllers

import (
	"bookLibrary/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (controller *BookController) GetBook(c *gin.Context) {
	passedParam := c.Param("id")

	id64, _ := strconv.ParseUint(passedParam, 10, 32)
	id := uint(id64)
	book, err := controller.bookService.GetBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func (controller *BookController) GetBooks(c *gin.Context) {
	books, err := controller.bookService.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, books)
	}
}

func (controller *BookController) UpdateBook(c *gin.Context) {
	var passedParam passedBook

	if err := c.BindJSON(&passedParam); err != nil {
		return
	}
	passedId := c.Param("id")
	id64, _ := strconv.ParseUint(passedId, 10, 32)
	id := uint(id64)

	book, err := controller.bookService.UpdateBook(id, passedParam.Name, passedParam.Author)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func (controller *BookController) DeleteBook(c *gin.Context) {
	passedParam := c.Param("id")

	id64, _ := strconv.ParseUint(passedParam, 10, 32)
	id := uint(id64)

	book, err := controller.bookService.DeleteBook(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, book)
	}
}

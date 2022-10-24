package services

import (
	"bookLibrary/models"
	"bookLibrary/repository"
)

type BookServiceInterface interface {
	AddBook(name string, author string) (uint, error)
	GetBook(id uint) (models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(id uint, name string, author string) (models.Book, error)
	DeleteBook(id uint) (models.Book, error)
}

type BookService struct {
	bookStorage repository.BookStorageInterface
}

func NewBookService(bookStorage repository.BookStorageInterface) *BookService {
	return &BookService{bookStorage: bookStorage}
}

func (service *BookService) AddBook(name string, author string) (uint, error) {

	return service.bookStorage.SaveBook(name, author)
}

func (service *BookService) GetBook(id uint) (models.Book, error) {
	return service.bookStorage.GetBook(id)
}

func (service *BookService) GetBooks() ([]models.Book, error) {
	return service.bookStorage.GetBooks()
}

func (service *BookService) UpdateBook(id uint, name string, author string) (models.Book, error) {
	return service.bookStorage.UpdateBook(id, name, author)
}

func (service *BookService) DeleteBook(id uint) (models.Book, error) {
	return service.bookStorage.DeleteBook(id)
}

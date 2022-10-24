package services

import "bookLibrary/repository"

type BookServiceInterface interface {
	AddBook(name string, author string) (uint, error)
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

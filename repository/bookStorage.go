package repository

import (
	"bookLibrary/config"
	"bookLibrary/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type BookStorageInterface interface {
	SaveBook(name string, author string) (uint, error)
	GetBook(id uint) (models.Book, error)
	GetBooks() ([]models.Book, error)
	UpdateBook(id uint, name string, author string) (models.Book, error)
	DeleteBook(id uint) (models.Book, error)
}

type BookStorage struct {
}

func (storage *BookStorage) SaveBook(name string, author string) (uint, error) {
	db := connectToDb()
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		return 0, err
	}

	book := &models.Book{
		Name:   name,
		Author: author,
	}
	_, err = db.Model(book).Insert()
	if err != nil {
		return 0, err
	}

	return book.Id, nil

}

func (storage *BookStorage) GetBook(id uint) (models.Book, error) {
	db := connectToDb()
	defer db.Close()

	book := &models.Book{Id: id}
	err := db.Model(book).WherePK().Select()

	return *book, err
}

func (storage *BookStorage) GetBooks() ([]models.Book, error) {
	db := connectToDb()
	defer db.Close()

	var books []models.Book
	err := db.Model(&books).Select()
	if err != nil {
		return nil, err
	}

	return books, err
}

func (storage *BookStorage) UpdateBook(id uint, name string, author string) (models.Book, error) {
	db := connectToDb()
	defer db.Close()
	book := &models.Book{
		Id:     id,
		Name:   name,
		Author: author,
	}
	_, err := db.Model(book).WherePK().Update(book)

	return *book, err
}

func (storage *BookStorage) DeleteBook(id uint) (models.Book, error) {
	db := connectToDb()
	defer db.Close()

	book := &models.Book{Id: id}
	_, err := db.Model(book).WherePK().Delete()

	return *book, err

}

func connectToDb() *pg.DB {
	cfg := config.Get()
	db := pg.Connect(&pg.Options{
		User:     cfg.PgUser,
		Password: cfg.PgPassword,
		Addr:     cfg.PgAddr,
		Database: cfg.PgDatabase,
	})
	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.Book)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}
	return nil
}

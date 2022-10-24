package repository

import (
	"bookLibrary/config"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type BookStorageInterface interface {
	SaveBook(name string, author string) (uint, error)
}

type BookStorage struct {
}

type Book struct {
	Id     uint
	Name   string
	Author string
}

func (storage *BookStorage) SaveBook(name string, author string) (uint, error) {
	db := connectToDb()
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		return 0, err
	}

	book := &Book{
		Name:   name,
		Author: author,
	}
	_, err = db.Model(book).Insert()
	if err != nil {
		return 0, err
	}

	return book.Id, nil

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
		(*Book)(nil),
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

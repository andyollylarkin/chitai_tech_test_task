package services

import (
	"github.com/andyollylarkin/chitai_tech_test_task/books"
	"github.com/andyollylarkin/chitai_tech_test_task/books/models"
	"github.com/rs/zerolog"
)

type BookService struct {
	repository books.Repository
	logger     zerolog.Logger
}

func NewBookService(repository books.Repository, logger zerolog.Logger) *BookService {
	return &BookService{repository: repository}
}

func (b *BookService) RetrieveBooks() []*models.Book {
	b.logger.Info().Msg("start retrieve books")
	return b.repository.GetAllBooks()
}

func (b *BookService) RemoveBook(id int64) error {
	err := b.repository.DeleteBookById(id)
	if err != nil {
		b.logger.Error().Err(err)
		return err
	}
	return nil
}

package books

import "github.com/andyollylarkin/chitai_tech_test_task/books/models"

type Repository interface {
	// GetAllBooks retrieve books from persistence storage
	GetAllBooks() []*models.Book
	// DeleteBookById delete book from storage
	DeleteBookById(id int64) error
}

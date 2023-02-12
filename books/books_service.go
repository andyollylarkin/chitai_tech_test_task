package books

import "github.com/andyollylarkin/chitai_tech_test_task/books/models"

type Service interface {
	RetrieveBooks() []*models.Book
	RemoveBook(id int64) error
}

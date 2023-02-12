package persistance

import (
	"errors"
	"github.com/andyollylarkin/chitai_tech_test_task/books/models"
	"sync"
)

type InMemoryRepository struct {
	collection []*models.Book
	mu         *sync.Mutex
}

// Репозиторий является синглтоном чтобы все запросы работали с одной и той же коллекцией
var instance *InMemoryRepository = nil

func NewInMemoryRepository() *InMemoryRepository {
	if instance == nil {
		var once sync.Once
		once.Do(func() {
			instance = &InMemoryRepository{collection: make([]*models.Book, 0), mu: new(sync.Mutex)}
		})
		return instance
	}
	return instance
}

// FillBooks fill books collection
// т.к. в задании не предусмотрено эндпоинта для добавления книг в коллекцию, то заполнять их будет все разом в коде
func (i *InMemoryRepository) FillBooks(booksCollection []*models.Book) {
	i.collection = booksCollection
}

// GetAllBooks retrieve books from persistence storage
func (i *InMemoryRepository) GetAllBooks() []*models.Book {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.collection
}

// DeleteBookById delete book from storage
func (i *InMemoryRepository) DeleteBookById(id int64) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	for idx, v := range i.collection {
		if v.GetId() == id {
			i.collection = append(i.collection[:idx], i.collection[idx+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

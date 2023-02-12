package models

import (
	"errors"
	"time"
)

var (
	errorChangeTitle  = errors.New("can't change title. New title can't be empty")
	errorChangeAuthor = errors.New("can't change author. New author can't be empty")
)

type Book struct {
	id            int64 // по хорошему в реальном сервисе я бы использовал UUID для идентификации сущности, но решил не усложнять
	title         string
	author        string
	publisherYear time.Time // храним год публикации в типе Time.
	// Если в будущем требования изменятся и нужен будет не только год, но и остальные части дат,
	//то у нас будет эта информация.
}

func NewBook(id int64, title string, author string, publisherYear time.Time) (*Book, error) {
	if title == "" || author == "" {
		return nil, errors.New("incorrect information for creating new book")
	}
	return &Book{id: id, title: title, author: author, publisherYear: publisherYear}, nil
}

func (b *Book) GetId() int64 {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetPublisherYear() int {
	return b.publisherYear.Year()
}

func (b *Book) ChangePublisherYear(toTime time.Time) {
	b.publisherYear = toTime
}

func (b *Book) Rename(newName string) error {
	if newName == "" {
		return errorChangeTitle
	}
	b.title = newName
	return nil
}

func (b *Book) ChangeAuthor(newAuthor string) error {
	if newAuthor == "" {
		return errorChangeAuthor
	}
	b.author = newAuthor
	return nil
}

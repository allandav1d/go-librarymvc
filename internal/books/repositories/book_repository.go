package repositories

import (
	"errors"
	"librarymvc/internal/books/models"
	"sync"
)

type BookRepository struct {
	books  map[int64]*models.Book
	mu     sync.RWMutex
	nextID int64
}

func NewBookRepository() models.BookRepository {
	return &BookRepository{
		books:  make(map[int64]*models.Book),
		nextID: 1,
	}
}

func (b *BookRepository) CreateBook(book *models.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	book.ID = b.nextID
	b.nextID++
	b.books[book.ID] = book

	return nil
}

func (b *BookRepository) GetBook(id int64) (*models.Book, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	book, exists := b.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}

	return book, nil
}

func (b *BookRepository) GetAllBooks() ([]*models.Book, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	books := make([]*models.Book, 0, len(b.books))
	for _, book := range b.books {
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepository) UpdateBook(id int64, book *models.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	_, exists := b.books[id]
	if !exists {
		return errors.New("book not found")
	}

	book.ID = id
	b.books[id] = book
	return nil
}

func (b *BookRepository) DeleteBook(id int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	_, exists := b.books[id]
	if !exists {
		return errors.New("book not found")
	}

	delete(b.books, id)
	return nil
}

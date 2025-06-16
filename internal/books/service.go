package books

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type BookService interface {
	CreateBook(book Book) (Book, error)
	GetAllBooks() ([]Book, error)
	GetBookByID(id int) (Book, error)
	UpdateBook(book Book) (Book, error)
	DeleteBook(id int) (map[string]string, error)
}

type bookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(book Book) (Book, error) {
	if book.PublishedDate == "" {
		book.PublishedDate = time.Now().Format("2006-01-02")
	}

	if _, err := time.Parse("2006-01-02", book.PublishedDate); err != nil {
		return Book{}, fmt.Errorf("invalid date format: must be YYYY-MM-DD")
	}

	return s.repo.Create(book)
}

func (s *bookService) GetAllBooks() ([]Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) GetBookByID(id int) (Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return Book{}, errors.New("book not found")
	}

	return book, nil
}

func (s *bookService) UpdateBook(book Book) (Book, error) {
	updatedBook, err := s.repo.Update(book)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Book{}, errors.New("book not found")
		}

		return Book{}, err
	}

	return updatedBook, nil
}

func (s *bookService) DeleteBook(id int) (map[string]string, error) {
	err := s.repo.Delete(id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("book not found")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to delete book: %w", err)
	}

	return map[string]string{"message": "book deleted successfully"}, nil
}

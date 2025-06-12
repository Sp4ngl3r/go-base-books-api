package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Sp4ngl3r/go-base-books-api/model"
	"github.com/Sp4ngl3r/go-base-books-api/repository"
)

type BookService interface {
	CreateBook(book model.Book) (model.Book, error)
	GetAllBooks() ([]model.Book, error)
	GetBookByID(id int) (model.Book, error)
	UpdateBook(book model.Book) (model.Book, error)
	DeleteBook(id int) (map[string]string, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(book model.Book) (model.Book, error) {
	return s.repo.Create(book)
}

func (s *bookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) GetBookByID(id int) (model.Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return model.Book{}, errors.New("book not found")
	}

	return book, nil
}

func (s *bookService) UpdateBook(book model.Book) (model.Book, error) {
	updatedBook, err := s.repo.Update(book)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, errors.New("book not found")
		}

		return model.Book{}, err
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

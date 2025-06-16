package books

import (
	"database/sql"
)

type BookRepository interface {
	Create(book Book) (Book, error)
	GetAll() ([]Book, error)
	GetByID(id int) (Book, error)
	Update(book Book) (Book, error)
	Delete(id int) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book Book) (Book, error) {
	query := `INSERT INTO books (title, author, published_date) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, book.Title, book.Author, book.PublishedDate).Scan(&book.ID)

	return book, err
}

func (r *bookRepository) GetAll() ([]Book, error) {
	query := `SELECT id, title, author, published_date FROM books`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedDate); err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	return books, nil
}

func (r *bookRepository) GetByID(id int) (Book, error) {
	query := `SELECT id, title, author, published_date FROM books WHERE id = $1`
	var b Book
	err := r.db.QueryRow(query, id).Scan(&b.ID, &b.Title, &b.Author, &b.PublishedDate)

	return b, err
}

func (r *bookRepository) Update(book Book) (Book, error) {
	query := `UPDATE books SET title=$1, author=$2, published_date=$3 WHERE id=$4`
	res, err := r.db.Exec(query, book.Title, book.Author, book.PublishedDate, book.ID)
	if err != nil {
		return Book{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return Book{}, err
	}
	if rowsAffected == 0 {
		return Book{}, sql.ErrNoRows
	}

	return book, nil
}

func (r *bookRepository) Delete(id int) error {
	query := `DELETE FROM books WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

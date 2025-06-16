package books

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	net_http "net/http"
	"strconv"

	"github.com/Sp4ngl3r/go-base-books-api/config"

	"github.com/unbxd/go-base/v2/log"
	"github.com/unbxd/go-base/v2/transport/http"
)

type BookHandler struct {
	bookService BookService
}

func NewBookHandler(service BookService) *BookHandler {
	return &BookHandler{bookService: service}
}

func (h *BookHandler) DecodeBook(_ context.Context, r *net_http.Request) (interface{}, error) {
	var b Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return nil, err
	}

	return b, nil
}

func (h *BookHandler) DecodeID(_ context.Context, r *net_http.Request) (interface{}, error) {
	idStr := http.Parameters(r).ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	return id, nil
}

func (h *BookHandler) DecodeBookWithID(ctx context.Context, r *net_http.Request) (interface{}, error) {
	var b Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return nil, err
	}

	idStr := http.Parameters(r).ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	b.ID = id

	return b, nil
}

func (h *BookHandler) EncodeResponse(_ context.Context, w net_http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(resp)
}

func (h *BookHandler) ErrorEncoder(ctx context.Context, err error, w net_http.ResponseWriter) {
	w.WriteHeader(net_http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}

func (h *BookHandler) Create(_ context.Context, req interface{}) (interface{}, error) {
	book := req.(Book)
	config.AppLogger.Info("creating book", log.String("title", book.Title), log.String("author", book.Author))

	createdBook, err := h.bookService.CreateBook(book)
	if err != nil {
		config.AppLogger.Error("failed to create book", log.String("title", book.Title), log.String("author", book.Author), log.Error(err))
		return nil, err
	}

	return createdBook, nil
}

func (h *BookHandler) GetAll(_ context.Context, _ interface{}) (interface{}, error) {
	config.AppLogger.Info("retrieving all books")

	return h.bookService.GetAllBooks()
}

func (h *BookHandler) Get(_ context.Context, req interface{}) (interface{}, error) {
	id := req.(int)
	config.AppLogger.Info("retrieving book with ID", log.Int("id", id))

	return h.bookService.GetBookByID(id)
}

func (h *BookHandler) Update(_ context.Context, req interface{}) (interface{}, error) {
	book := req.(Book)

	updatedBook, err := h.bookService.UpdateBook(book)
	if err != nil {
		config.AppLogger.Error("failed to update book", log.Int("id", book.ID), log.Error(err))
		return nil, err
	}

	config.AppLogger.Info("book updated successfully", log.Int("id", book.ID), log.String("title", book.Title), log.String("author", book.Author))

	return updatedBook, nil
}

func (h *BookHandler) Delete(_ context.Context, req interface{}) (interface{}, error) {
	id := req.(int)

	resp, err := h.bookService.DeleteBook(id)
	if err != nil {
		if err.Error() == "book not found" {
			config.AppLogger.Error("book not found during deletion", log.Int("id", id))
			return nil, err
		}

		config.AppLogger.Error("failed to delete book", log.Int("id", id), log.Error(err))
		return nil, err
	}

	config.AppLogger.Info("book deleted successfully", log.Int("id", id))

	return resp, nil
}

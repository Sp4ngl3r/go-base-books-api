package transport

import (
	"github.com/Sp4ngl3r/go-base-books-api/handler"
	"github.com/unbxd/go-base/v2/transport/http"
)

func SetupRoutes(h *handler.BookHandler) (*http.Transport, error) {
	tr, err := http.NewHTTPTransport(
		"book-api",
		http.WithCustomHostPort("0.0.0.0", "5555"),
		http.WithTransportOption(http.WithErrorEncoder(h.ErrorEncoder)),
	)
	if err != nil {
		return nil, err
	}

	tr.GET("/api/v1/books", h.GetAll, http.HandlerWithEncoder(h.EncodeResponse))
	tr.GET("/api/v1/books/{id}", h.Get, http.HandlerWithDecoder(h.DecodeID), http.HandlerWithEncoder(h.EncodeResponse))
	tr.POST("/api/v1/books", h.Create, http.HandlerWithDecoder(h.DecodeBook), http.HandlerWithEncoder(h.EncodeResponse))
	tr.PUT("/api/v1/books/{id}", h.Update, http.HandlerWithDecoder(h.DecodeBookWithID), http.HandlerWithEncoder(h.EncodeResponse))
	tr.DELETE("/api/v1/books/{id}", h.Delete, http.HandlerWithDecoder(h.DecodeID), http.HandlerWithEncoder(h.EncodeResponse))

	return tr, nil
}

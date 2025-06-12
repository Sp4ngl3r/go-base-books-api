package main

import (
	"github.com/Sp4ngl3r/go-base-books-api/config"
	"github.com/Sp4ngl3r/go-base-books-api/handler"
	"github.com/Sp4ngl3r/go-base-books-api/repository"
	"github.com/Sp4ngl3r/go-base-books-api/service"
	"github.com/Sp4ngl3r/go-base-books-api/transport"
	"github.com/unbxd/go-base/v2/log"
)

func main() {

	config.InitLogger()
	config.LoadConfig()
	defer config.AppConfig.DB.Close()

	bookRepo := repository.NewBookRepository(config.AppConfig.DB)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	tr, err := transport.SetupRoutes(bookHandler)
	if err != nil {
		config.AppLogger.Fatal("Failed to setup routes: %v", log.Error(err))
	}

	config.AppLogger.Info("🚀 Starting server...",
		log.String("port", config.AppConfig.AppPort),
	)

	tr.Open()
}

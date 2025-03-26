package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/config"
	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/db"
	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/http"
	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/http/handlers"
	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/logger"
	usecases "github.com/atcheri/warehouse-api-go-tdd/internal/use-cases"
)

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init store
	inMemoryProductStore := db.NewInMemoryDB()

	// Init router
	router, err := http.NewRouter(
		config.HTTP,
		handlers.NewHelloHandler(),
		handlers.NewProductHandler(usecases.NewCreateProductUsecase(inMemoryProductStore)),
	)

	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}

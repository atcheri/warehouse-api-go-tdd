package doubles

import (
	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/config"
)

func NewTestConfig() (*config.Container, error) {
	app := &config.App{
		Name: "warehouse-tdd-test",
		Env:  "test",
	}

	db := &config.DB{}

	http := &config.HTTP{
		Env:            "test",
		URL:            "http://127.0.0.1",
		Port:           "9999",
		AllowedOrigins: "*",
	}

	return &config.Container{
		App:  app,
		DB:   db,
		HTTP: http,
	}, nil
}

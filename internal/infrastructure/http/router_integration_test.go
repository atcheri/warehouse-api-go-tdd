package http_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/db"
	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/doubles"
	rest "github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/http"
	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/http/handlers"
	usecases "github.com/atcheri/warehouse-api-go-tdd/internal/use-cases"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Run("POST request to the product endpoint stores a new product in the warehouse", func(t *testing.T) {
		// arrange
		config, _ := doubles.NewTestConfig()
		productHandler := handlers.NewProductHandler(usecases.NewCreateProductUsecase(db.NewInMemoryDB()))
		server, _ := rest.NewRouter(config.HTTP, handlers.NewHelloHandler(), productHandler)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/v1/product", bytes.NewBufferString(`{
			"name": "product name",
			"price": 13.45
		}`,
		))

		// act
		server.ServeHTTP(w, req)
		result := w.Result()
		id := result.Header.Get("id")

		// assert
		assert.Equal(t, http.StatusCreated, result.StatusCode)
		assert.NotEmpty(t, id)
	})

	t.Run("fails when the product payload doesn't have a name", func(t *testing.T) {
		// arrange
		config, _ := doubles.NewTestConfig()
		productHandler := handlers.NewProductHandler(usecases.NewCreateProductUsecase(db.NewInMemoryDB()))
		server, _ := rest.NewRouter(config.HTTP, handlers.NewHelloHandler(), productHandler)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/v1/product", bytes.NewBuffer([]byte(`{
			"price": 13.45
		}`),
		))

		// act
		server.ServeHTTP(w, req)
		result := w.Result()

		// assert
		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})

	t.Run("fails when the product payload has an empty string", func(t *testing.T) {
		// arrange
		config, _ := doubles.NewTestConfig()
		productHandler := handlers.NewProductHandler(usecases.NewCreateProductUsecase(db.NewInMemoryDB()))
		server, _ := rest.NewRouter(config.HTTP, handlers.NewHelloHandler(), productHandler)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/v1/product", bytes.NewBuffer([]byte(`{
			"name": "",
			"price": 13.45
		}`),
		))

		// act
		server.ServeHTTP(w, req)
		result := w.Result()

		// assert
		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})

	t.Run("fails when trying to create a product with the same name", func(t *testing.T) {
		// arrange
		config, _ := doubles.NewTestConfig()
		productHandler := handlers.NewProductHandler(usecases.NewCreateProductUsecase(db.NewInMemoryDB()))
		server, _ := rest.NewRouter(config.HTTP, handlers.NewHelloHandler(), productHandler)
		w := httptest.NewRecorder()
		req1, _ := http.NewRequest(http.MethodPost, "/v1/product", bytes.NewBuffer([]byte(`{
			"name": "duplicate",
			"price": 13.45
		}`),
		))
		req2, _ := http.NewRequest(http.MethodPost, "/v1/product", bytes.NewBuffer([]byte(`{
			"name": "duplicate",
			"price": 13.45
		}`),
		))

		// act
		server.ServeHTTP(httptest.NewRecorder(), req1)
		server.ServeHTTP(w, req2)
		result := w.Result()

		// assert
		assert.Equal(t, http.StatusConflict, result.StatusCode)
	})
}

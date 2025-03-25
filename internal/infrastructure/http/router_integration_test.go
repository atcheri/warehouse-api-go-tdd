package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/doubles"
	rest "github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/http"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Run("POST request to the product endpoint stores a new product in the warehouse", func(t *testing.T) {
		// arrange
		config, _ := doubles.NewTestConfig()
		server, _ := rest.NewRouter(config.HTTP, rest.NewHelloHandler(), rest.NewProductHandler())
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/v1/product", nil)

		// act
		server.ServeHTTP(w, req)

		// assert
		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
	})
}

package http

import (
	"log/slog"
	"strings"

	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/config"
	"github.com/atcheri/warehouse-api-go-tdd/internal/infrastructure/http/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	config *config.HTTP,
	helloHandler *handlers.HelloHandler,
	productHandler *handlers.ProductHandler,
) (*Router, error) {
	// Disable debug mode in production
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	// Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		hello := v1.Group("/hello")
		{
			hello.GET("", helloHandler.HelloWorld)
		}

		product := v1.Group("/product")
		{
			product.POST("", productHandler.CreateProduct)
			product.GET("/:id", productHandler.RetrieveProduct)
		}
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

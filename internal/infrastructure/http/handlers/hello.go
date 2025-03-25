package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) HelloWorld(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world")
}

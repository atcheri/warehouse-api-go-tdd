package handlers

import (
	"net/http"

	usecases "github.com/atcheri/warehouse-api-go-tdd/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required,min=1" example:"New Product"`
	Price float64 `json:"price" binding:"required,number" example:"5.69"`
}

type ProductHandler struct {
	create usecases.CreateProduct
}

func NewProductHandler(create usecases.CreateProduct) *ProductHandler {
	return &ProductHandler{create}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var createProductRequest CreateProductRequest
	err := ctx.ShouldBindJSON(&createProductRequest)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, _ := h.create.Execute(createProductRequest.Name, createProductRequest.Price)
	ctx.Header("id", product.ID.String())
	ctx.JSON(http.StatusCreated, nil)
}

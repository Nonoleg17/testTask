package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testCase/internal/entity"
	"testCase/internal/usecase"
	"testCase/pkg/logger"
)

type ProductCreateRequest struct {
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
	LeftInStock int    `json:"left_in_stock"`
}
type ProductDeleteRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
type ProductUpdateRequest struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	Currency    string    `json:"currency" binding:"required"`
	LeftInStock int       `json:"left_in_stock"`
}
type ProductRoutes struct {
	u usecase.Product
	l logger.Interface
}

func newProductRoutes(handler *gin.RouterGroup, u usecase.Product, l logger.Interface) {
	r := &ProductRoutes{u, l}
	h := handler.Group("/product")
	{
		h.POST("/create-product", r.CreateProduct)
		h.DELETE("/delete-product", r.DeleteProduct)
		h.PATCH("/update-product", r.UpdateProduct)
	}
}

func (r *ProductRoutes) CreateProduct(c *gin.Context) {
	var request ProductCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - CreateProduct")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.CreateProduct(c.Request.Context(), &entity.Product{
		Description: request.Description,
		Price:       request.Price,
		Currency:    request.Currency,
		LeftInStock: request.LeftInStock,
	})
	if err != nil {
		r.l.Error(err, "http - CreateProduct")
		errorResponse(c, http.StatusInternalServerError, "create Product problem")
		return
	}
	c.JSON(http.StatusOK, "Success create")
}

func (r *ProductRoutes) DeleteProduct(c *gin.Context) {
	var request ProductDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - DeleteProduct")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.DeleteProduct(c.Request.Context(), request.Id)
	if err != nil {
		r.l.Error(err, "http - DeleteProduct")
		errorResponse(c, http.StatusInternalServerError, "delete Product problem")
		return
	}
	c.JSON(http.StatusOK, "Success delete")
}

func (r *ProductRoutes) UpdateProduct(c *gin.Context) {
	var request ProductUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - UpdateProduct")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.UpdateProduct(c.Request.Context(), &entity.Product{
		ID:          request.Id,
		Description: request.Description,
		Price:       request.Price,
		Currency:    request.Currency,
		LeftInStock: request.LeftInStock,
	})
	if err != nil {
		r.l.Error(err, "http - UpdateProduct")
		errorResponse(c, http.StatusInternalServerError, "update Product problem")
		return
	}
	c.JSON(http.StatusOK, "Success update")

}

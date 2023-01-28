package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testCase/internal/entity"
	"testCase/internal/usecase"
	"testCase/pkg/logger"
)

type OrderCreateRequest struct {
	UserId   uuid.UUID  `json:"user_id" binding:"required"`
	Products []Products `json:"products" binding:"required"`
}
type Products struct {
	Id    uuid.UUID `json:"id" binding:"required"`
	Count int       `json:"count" binding:"required"`
}
type OrderDeleteRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
type OrderUpdateRequest struct {
	Id       uuid.UUID  `json:"id" binding:"required"`
	UserId   uuid.UUID  `json:"user_id" binding:"required"`
	Products []Products `json:"products" binding:"required"`
}
type orderRoutes struct {
	u   usecase.Order
	uOP usecase.OrderProduct
	l   logger.Interface
}

func newOrderRoutes(handler *gin.RouterGroup, u usecase.Order, uOP usecase.OrderProduct, l logger.Interface) {
	r := &orderRoutes{u, uOP, l}
	h := handler.Group("/order")
	{
		h.POST("/create-order", r.CreateOrder)
		h.DELETE("/delete-order", r.DeleteOrder)
		h.PATCH("/update-order", r.UpdateOrder)
	}
}

func (r *orderRoutes) CreateOrder(c *gin.Context) {
	var request OrderCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - CreateOrder")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	orderID, err := r.u.CreateOrder(c.Request.Context(), &entity.Order{
		UserId: request.UserId,
	})
	if err != nil {
		r.l.Error(err, "http - CreateOrder")
		errorResponse(c, http.StatusInternalServerError, "create Order problem")
		return
	}
	for _, value := range request.Products {

		err = r.uOP.CreateOrderProduct(c.Request.Context(),
			&entity.OrderProduct{OrderId: orderID, ProductId: value.Id, CountProducts: value.Count})
		if err != nil {
			r.l.Error(err, "http - CreateOrder")
		}
	}
	if err != nil {
		r.l.Error(err, "http - CreateOrder")
		errorResponse(c, http.StatusInternalServerError, "create Order problem")
		return
	}

	c.JSON(http.StatusOK, "Success create")
}

func (r *orderRoutes) DeleteOrder(c *gin.Context) {
	var request OrderDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - DeleteOrder")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.DeleteOrder(c.Request.Context(), request.Id)
	if err != nil {
		r.l.Error(err, "http - DeleteOrder")
		errorResponse(c, http.StatusInternalServerError, "delete Order problem")
		return
	}
	c.JSON(http.StatusOK, "Success delete")
}

func (r *orderRoutes) UpdateOrder(c *gin.Context) {
	var request OrderUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - UpdateOrder")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.UpdateOrder(c.Request.Context(), &entity.Order{
		ID:     request.Id,
		UserId: request.UserId,
	})
	if err != nil {
		r.l.Error(err, "http - UpdateOrder")
		errorResponse(c, http.StatusInternalServerError, "update order Order problem")
		return
	}
	err = r.uOP.ClearAllOrderProduct(c.Request.Context(), request.Id)
	if err != nil {
		r.l.Error(err, "http - UpdateOrder")
		errorResponse(c, http.StatusInternalServerError, "update OrderProduct problem")
		return
	}
	for _, value := range request.Products {
		err = r.uOP.CreateOrderProduct(c.Request.Context(),
			&entity.OrderProduct{OrderId: request.Id, ProductId: value.Id, CountProducts: value.Count})
		if err != nil {
			r.l.Error(err, "http - CreateOrder")
		}
	}
	c.JSON(http.StatusOK, "Success update")

}

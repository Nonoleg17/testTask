package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testCase/internal/entity"
	"testCase/internal/usecase"
	"testCase/pkg/logger"
)

type UserRequest struct {
	Firstname  string `json:"firstname" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Middlename string `json:"middlename" binding:"required"`
	Sex        string `json:"sex" binding:"required"`
	Age        int    `json:"age" binding:"required"`
}

type userRoutes struct {
	u usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.User, l logger.Interface) {
	r := &userRoutes{u, l}
	h := handler.Group("/user")
	{
		h.POST("/create-user", r.CreateUser)
	}
}

func (r *userRoutes) CreateUser(c *gin.Context) {
	var request UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - CreateUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.CreateUser(c.Request.Context(), &entity.User{
		Firstname:  request.Firstname,
		Surname:    request.Surname,
		Middlename: request.Middlename,
		Sex:        request.Sex,
		Age:        request.Age,
	})
	if err != nil {
		r.l.Error(err, "http - CreateUser")
		errorResponse(c, http.StatusInternalServerError, "create user problem")
		return
	}
	c.JSON(http.StatusOK, "Success")
}

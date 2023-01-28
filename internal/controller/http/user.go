package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testCase/internal/entity"
	"testCase/internal/usecase"
	"testCase/pkg/logger"
)

type UserCreateRequest struct {
	Firstname  string `json:"firstname" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Middlename string `json:"middlename" binding:"required"`
	Sex        string `json:"sex" binding:"required"`
	Age        int    `json:"age" binding:"required"`
}
type UserDeleteRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
type UserUpdateRequest struct {
	Id         uuid.UUID `json:"id" binding:"required"`
	Firstname  string    `json:"firstname" binding:"required"`
	Surname    string    `json:"surname" binding:"required"`
	Middlename string    `json:"middlename" binding:"required"`
	Sex        string    `json:"sex" binding:"required"`
	Age        int       `json:"age" binding:"required"`
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
		h.DELETE("/delete-user", r.DeleteUser)
		h.PATCH("/update-user", r.UpdateUser)
	}
}

func (r *userRoutes) CreateUser(c *gin.Context) {
	var request UserCreateRequest
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
	c.JSON(http.StatusOK, "Success create")
}

func (r *userRoutes) DeleteUser(c *gin.Context) {
	var request UserDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - DeleteUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.DeleteUser(c.Request.Context(), request.Id)
	if err != nil {
		r.l.Error(err, "http - DeleteUser")
		errorResponse(c, http.StatusInternalServerError, "delete user problem")
		return
	}
	c.JSON(http.StatusOK, "Success delete")
}

func (r *userRoutes) UpdateUser(c *gin.Context) {
	var request UserUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - UpdateUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.UpdateUser(c.Request.Context(), &entity.User{
		ID:         request.Id,
		Firstname:  request.Firstname,
		Surname:    request.Surname,
		Middlename: request.Middlename,
		Sex:        request.Sex,
		Age:        request.Age,
	})
	if err != nil {
		r.l.Error(err, "http - UpdateUser")
		errorResponse(c, http.StatusInternalServerError, "update user problem")
		return
	}
	c.JSON(http.StatusOK, "Success update")

}

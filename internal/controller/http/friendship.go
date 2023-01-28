package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testCase/internal/usecase"
	"testCase/pkg/logger"
)

type friendshipRequest struct {
	FirstUserId  uuid.UUID `json:"first_user_id"  binding:"required"`
	SecondUserId uuid.UUID `json:"second_user_id"  binding:"required"`
}

type friendshipRoutes struct {
	u usecase.Friendship
	l logger.Interface
}

func newFriendshipRoutes(handler *gin.RouterGroup, u usecase.Friendship, l logger.Interface) {
	r := &friendshipRoutes{u, l}
	h := handler.Group("/friendship")
	{
		h.POST("/follow", r.Follow)
	}
}

func (r *friendshipRoutes) Follow(c *gin.Context) {
	var request friendshipRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - Follow")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	err := r.u.FollowUser(c.Request.Context(), request.FirstUserId, request.SecondUserId)
	if err != nil {
		r.l.Error(err, "http - Follow")
		errorResponse(c, http.StatusInternalServerError, "Follow user problem")
		return
	}
	c.JSON(http.StatusOK, "Success follow")
}

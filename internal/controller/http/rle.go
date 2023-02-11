package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testCase/internal/usecase"
	"testCase/pkg/logger"
)

type EncodeDecodeRequest struct {
	Text []string `json:"text" binding:"required"`
}

type RleRoutes struct {
	u usecase.Rle
	l logger.Interface
}

func newRleRoutes(handler *gin.RouterGroup, u usecase.Rle, l logger.Interface) {
	r := &RleRoutes{u, l}
	h := handler.Group("/rle")
	{
		h.POST("/encode", r.Encode)
		h.POST("/decode", r.Decode)
	}
}

func (r *RleRoutes) Encode(c *gin.Context) {
	var request EncodeDecodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - Encode")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	res, err := r.u.RunLengthEncode(c.Request.Context(), request.Text)
	if err != nil {
		r.l.Error(err, "http - Encode")
		errorResponse(c, http.StatusInternalServerError, "Encode problem")
		return
	}
	c.JSON(http.StatusOK, res)
}
func (r *RleRoutes) Decode(c *gin.Context) {
	var request EncodeDecodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - Decode")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	res, err := r.u.RunLengthDecode(c.Request.Context(), request.Text)
	if err != nil {
		r.l.Error(err, "http - Decode")
		errorResponse(c, http.StatusInternalServerError, "Decode problem")
		return
	}
	c.JSON(http.StatusOK, res)
}

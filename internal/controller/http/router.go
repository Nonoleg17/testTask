package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"testCase/internal/usecase"
	"testCase/pkg/logger"
)

func NewRouter(handler *gin.Engine, l logger.Interface, u usecase.User, p usecase.Product, f usecase.Friendship) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)
	h := handler.Group("/")
	{
		newUserRoutes(h, u, l)
		newProductRoutes(h, p, l)
		newFriendshipRoutes(h, f, l)
	}

}

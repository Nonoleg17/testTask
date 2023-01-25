package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testCase/config"
	"testCase/pkg/httpserver"
	"testCase/pkg/logger"
	"testCase/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	//repository
	_, err := postgres.New(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	handler := gin.New()
	v1.NewRouter(handler, l, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTPPort))

}

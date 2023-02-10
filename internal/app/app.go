package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"testCase/config"
	"testCase/internal/controller/http"
	"testCase/internal/usecase"
	"testCase/internal/usecase/repo"
	"testCase/pkg/httpserver"
	"testCase/pkg/logger"
	"testCase/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	//repository
	pg, err := postgres.New(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	rleUsecase := usecase.NewRleUseCase(repo.NewRleRepo(pg))

	handler := gin.New()
	http.NewRouter(handler, l, rleUsecase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTPPort))
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}

package handlers

import (
	"context"
	"golang-project-template/logger"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	log = logger.GetLogger("handlers")
	e   = echo.New()
)

func startServer() {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	publicAPIGroup := e.Group("/api/v1/public")

	publicAPIGroup.POST("/ad", createAd)
	publicAPIGroup.GET("/ads", listAd)
	publicAPIGroup.GET("/ad", getAdDetail)
	publicAPIGroup.PUT("/ad", updateAd)
	publicAPIGroup.DELETE("/ad", deleteAd)

	log.Info("Starting at port: 8000")
	if err := e.Start(":8000"); err != nil {
		log.Error(err)
	}
}

func waitForInterruptSignal() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	// stopClients()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Error(err)
	}
}

// Start APIs service
func Start() {
	go startServer()
	waitForInterruptSignal()
}

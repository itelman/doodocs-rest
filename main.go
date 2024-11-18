package main

import (
	"log"
	"net/http"
	"time"

	"github.com/itelman/doodocs-rest/config"
	"github.com/itelman/doodocs-rest/internal/handler"
	"github.com/itelman/doodocs-rest/internal/service"
	"github.com/itelman/doodocs-rest/pkg/logger"
)

func main() {
	config := config.NewConfig()
	service := service.NewService(config)
	delivery := handler.NewHandler(service)

	server := &http.Server{
		Addr:           config.Host + ":" + config.Port,
		Handler:        delivery.Routes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logger.InfoLog.Printf("Server run on http://localhost:%s", config.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error running server: %s\n", err)
	}
}

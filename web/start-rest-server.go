package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"one-pte-backend/config"
	"one-pte-backend/web/handlers"
	"one-pte-backend/web/swagger"
	"sync"
)

func StartRestServer(wg *sync.WaitGroup) {
	conf := config.GetConfig()
	router := gin.New()
	if conf.Mode == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// setup recovery middleware
	router.Use(gin.Recovery())

	// setup cors
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	// setup handlers here
	if conf.Mode != config.ReleaseMode {
		swagger.SetupSwagger(router)
	}

	if conf.HealthCheckRoute != "" {
		router.POST(conf.HealthCheckRoute, handlers.HealthCheck)
	}

	router.GET("/hello", handlers.Hello)

	wg.Add(1)

	go func() {
		slog.Info(fmt.Sprintf("HTTP Server Listening at %v", conf.HttpPort))

		defer wg.Done()

		if err := router.Run(fmt.Sprintf(":%d", conf.HttpPort)); err != nil {
			slog.Error(err.Error())
		}
	}()
}

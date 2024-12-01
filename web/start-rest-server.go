package web

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"one-pte-backend/config"
	"one-pte-backend/web/handlers"
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

	if conf.HealthCheckRoute != "" {
		router.POST(conf.HealthCheckRoute, handlers.HealthCheck)
	}

	router.GET("/hello", handlers.Hello)
	router.GET("/get-questions", handlers.GetQuestions)
	router.GET("/get-quetion-details", handlers.GetQuestionDetails)
	router.GET("/get-user-history", handlers.GetUserHistory)

	router.POST("/submit-answer", handlers.SubmitAnswer)

	wg.Add(1)

	go func() {
		slog.Info(fmt.Sprintf("HTTP Server Listening at %v", conf.HttpPort))

		defer wg.Done()

		if err := router.Run(fmt.Sprintf(":%d", conf.HttpPort)); err != nil {
			slog.Error(err.Error())
		}
	}()
}

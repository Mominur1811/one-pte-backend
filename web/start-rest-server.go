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

	if conf.Mode == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	router.GET("/hello", handlers.Hello)
	router.GET("/get-questions", handlers.GetQuestions)
	router.GET("/get-question-details", handlers.GetQuestionDetails)
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

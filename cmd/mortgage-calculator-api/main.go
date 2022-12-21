package main

import (
	"log"
	"mortgage-calculator-api/internal/app/service"

	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(cors.Default())
	router.POST("/calculate", service.CalcMortgage)
	router.GET("/ping", service.Ping)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

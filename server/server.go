package server

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/heyubani/go-template/config"
	"github.com/heyubani/go-template/core/base"
	v1 "github.com/heyubani/go-template/server/v1"
)

func InitHttpServer() {
	port := config.AppConfig.AppPort

	router := gin.Default()
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Cache-Control", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config),
		gin.Logger(),
		gin.Recovery(),
		requestid.New(requestid.WithCustomHeaderStrKey(base.REQUEST_ID_HEADER_NAME)),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Pings Active",
		})
	})

	router.GET("/health/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is healthy",
		})
	})

	v1.SetRoutes(router)

	log.Println("Started Application Server On Port: ", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting http server", err)
	}

}

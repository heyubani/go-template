package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heyubani/go-template/modules/routes"

	"golang.org/x/time/rate"
)

const ROUTE_GROUP = "/api/v1"

var limiter = rate.NewLimiter(1, 30)

func rateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many request"})
		c.Abort()
		return
	}
}

func SetRoutes(router *gin.Engine) {
	router.Use(rateLimiter)
	mainRouter := router.Group(ROUTE_GROUP)
	routes.SetModuleRoutes(*mainRouter)
}

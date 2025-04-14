package route

import (
	"github.com/gin-gonic/gin"
	"github.com/heyubani/go-template/core/container"
	"github.com/heyubani/go-template/modules/auth/controller"
)

const AUTH_BASE_URL = "/auth"

// this route is created for strictly for auth module,
// which means that you can create each route per the module

func RegisterAuthRoutes(router gin.RouterGroup) {
	authRoute := router.Group(AUTH_BASE_URL)

	authController := container.CreateControllerInvoker(controller.NewAuthController)
	{

		// for protected route
		// protectedRoute := salesRoute.Use(middleware.JwtMiddleware(), middleware.ServiceIDMiddleware(), middleware.AuditMiddleware())
		authRoute.GET("/", authController.Call("GetUser"))

	}
}

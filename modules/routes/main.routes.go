package routes

import (
	"github.com/gin-gonic/gin"
	auth "github.com/heyubani/go-template/modules/auth/route"
)

// This is we set all module route .. this is the global route group
func SetModuleRoutes(router gin.RouterGroup) {
	auth.RegisterAuthRoutes(router)
}

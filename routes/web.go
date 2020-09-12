package routes

import (
	ctrl "anwarmedika/hrm/controllers"

	"github.com/gin-gonic/gin"
)

// WebRoute : web routes
func WebRoute(r *gin.Engine) {

	user := new(ctrl.UserController)
	r.GET("/", user.Create)

	auth := new(ctrl.AuthController)
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/register", auth.Register)
		authRoute.POST("/logout", auth.Logout)
	}
}

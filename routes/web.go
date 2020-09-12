package routes

import (
	ctrl "anwarmedika/hrm/controllers"

	"github.com/gin-gonic/gin"
)

// WebRoute : web routes
func WebRoute(r *gin.Engine) {

	user := new(ctrl.UserController)
	r.GET("/", user.Create)
}

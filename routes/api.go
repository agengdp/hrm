package routes

import (
	"anwarmedika/hrm/controllers"

	"github.com/gin-gonic/gin"
)

// ApiRoute : web routes
func ApiRoute(r *gin.Engine) {

	health := new(controllers.HealthController)
	r.GET("/", health.Status)
}

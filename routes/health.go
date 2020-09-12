package routes

import (
	"anwarmedika/hrm/controllers"

	"github.com/gin-gonic/gin"
)

// HealthzRoute : web routes
func HealthzRoute(r *gin.Engine) {

	health := new(controllers.HealthController)
	r.GET("/healthz", health.Status)
}

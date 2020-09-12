package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController :
type HealthController struct{}

// Status :
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

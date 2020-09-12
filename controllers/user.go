package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// UserController :
type UserController struct{}

// Create :
func (ctrl UserController) Create(c *gin.Context) {
	c.String(http.StatusOK, os.Getenv("DB_PASS"))
}

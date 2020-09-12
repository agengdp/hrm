package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/twinj/uuid"
)

// RequestIDMidlleware : generate unique ID and attach it to each request for future reference or use
func RequestIDMidlleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

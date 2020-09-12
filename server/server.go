package server

import (
	mid "anwarmedika/hrm/middlewares"
	"anwarmedika/hrm/routes"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// Init : Start server
func Init() {

	r := gin.Default()
	r.Use(mid.CORSMiddleware())
	r.Use(mid.RequestIDMidlleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	routes.HealthzRoute(r)
	routes.ApiRoute(r)
	r.Run()

}

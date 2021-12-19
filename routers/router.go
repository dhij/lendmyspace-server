package routers

import (
	"github.com/davidhwang-ij/study-platform/internal/auth/http"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", http.Create)

	return r
}

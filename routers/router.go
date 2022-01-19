package routers

import (
	http2 "lendmyspace-server/internal/space/http"
	"lendmyspace-server/internal/user/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *http.UserHandler, spaceHandler *http2.SpaceHandler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/users/:user_id", userHandler.GetUser)
	r.GET("/users", userHandler.ListUsers)
	r.POST("/signup", userHandler.CreateUser)
	r.PATCH("/users/:user_id", userHandler.UpdateUser)
	r.DELETE("/users/:user_id", userHandler.DeleteUser)

	r.GET("/spaces/:space_id", spaceHandler.GetSpace)
	r.POST("/create_space", spaceHandler.CreateSpace)
}

func Start(serverAddress string) error {
	return r.Run(serverAddress)
}

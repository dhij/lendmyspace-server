package routers

import (
	http2 "lendmyspace-server/internal/space/http"
	"lendmyspace-server/internal/user/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *http.UserHandler, spaceHandler *http2.SpaceHandler) {
	r = gin.Default()

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

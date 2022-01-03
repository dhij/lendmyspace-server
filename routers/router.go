package routers

import (
	"dplatform/internal/user/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *http.UserHandler) {
	r = gin.Default()

	r.GET("/users/:user_id", userHandler.GetUser)
	r.GET("/users", userHandler.ListUsers)
	r.POST("/signup", userHandler.CreateUser)
	r.PATCH("/users/:user_id", userHandler.UpdateUser)
	r.DELETE("/users/:user_id", userHandler.DeleteUser)
}

func Start(serverAddress string) error {
	return r.Run(serverAddress)
}

package main

import (
	"dplatform/db"
	"dplatform/internal/user/http"
	"dplatform/internal/user/repository"
	"dplatform/internal/user/service"
	"dplatform/routers"
	"log"
)

func main() {
	r := routers.InitRouter()
	dbSQLX, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}
	defer dbSQLX.Close()

	userRepository := repository.NewUserRepository(dbSQLX.GetDB())
	userService := service.NewUserSerivce(userRepository)
	userHandler := http.NewUserHandler(userService)

	r.POST("/signup", userHandler.CreateUser)

	r.Run(":8080")
}

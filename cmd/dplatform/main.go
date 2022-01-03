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
	dbSQLX, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}
	defer dbSQLX.Close()

	userRepository := repository.NewUserRepository(dbSQLX.GetDB())
	userService := service.NewUserSerivce(userRepository)
	userHandler := http.NewUserHandler(userService)

	routers.InitRouter(userHandler)
	err = routers.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}

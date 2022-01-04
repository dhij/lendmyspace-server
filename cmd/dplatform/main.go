package main

import (
	"dplatform/db"
	http2 "dplatform/internal/room/http"
	repository2 "dplatform/internal/room/repository"
	service2 "dplatform/internal/room/service"
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

	roomRepository := repository2.NewRoomRepository(dbSQLX.GetDB())
	roomService := service2.NewRoomService(roomRepository)
	roomHandler := http2.NewRoomHandler(roomService)

	routers.InitRouter(userHandler, roomHandler)
	err = routers.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}

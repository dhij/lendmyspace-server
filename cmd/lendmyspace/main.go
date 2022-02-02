package main

import (
	"lendmyspace-server/db"
	http2 "lendmyspace-server/internal/space/http"
	repository2 "lendmyspace-server/internal/space/repository"
	service2 "lendmyspace-server/internal/space/service"
	"lendmyspace-server/internal/user/http"
	"lendmyspace-server/internal/user/repository"
	"lendmyspace-server/internal/user/service"
	"lendmyspace-server/routers"
	"lendmyspace-server/util"

	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbSQLX, err := db.NewDatabase(config)
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}
	defer dbSQLX.Close()

	userRepository := repository.NewUserRepository(dbSQLX.GetDB())
	userService := service.NewUserSerivce(userRepository)
	userHandler := http.NewUserHandler(userService)

	spaceRepository := repository2.NewSpaceRepository(dbSQLX.GetDB())
	spaceService := service2.NewSpaceService(spaceRepository)
	spaceHandler := http2.NewSpaceHandler(spaceService)

	routers.InitRouter(userHandler, spaceHandler)
	err = routers.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}

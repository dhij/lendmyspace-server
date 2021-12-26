package main

import (
	"dplatform/db"
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

	r.Run(":8080")
}

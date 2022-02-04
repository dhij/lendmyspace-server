package repository

import (
	"lendmyspace-server/db"
	"lendmyspace-server/util"
	"log"
	"os"
	"testing"
)

var dbSQLX *db.PostgreSQL

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbSQLX, err = db.NewDatabase(config)
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}

	os.Exit(m.Run())
}

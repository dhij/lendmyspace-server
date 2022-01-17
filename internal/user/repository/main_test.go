package repository

import (
	"lendmyspace-server/db"
	"log"
	"os"
	"testing"
)

var dbSQLX *db.PostgreSQL

func TestMain(m *testing.M) {
	var err error
	dbSQLX, err = db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}

	os.Exit(m.Run())
}

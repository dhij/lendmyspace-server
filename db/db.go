package db

import (
	"lendmyspace-server/util"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	maxIdleConns = 10
	maxOpenConns = 10
	maxConnLife  = 5 * time.Minute
)

type PostgreSQL struct {
	db *sqlx.DB
}

// NewDatabase make new database connection
func NewDatabase(config util.Config) (*PostgreSQL, error) {
	connectionStr := config.DBSource
	db, err := sqlx.Open(config.DBDriver, connectionStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(maxConnLife)

	return &PostgreSQL{
		db: db,
	}, nil
}

func (p *PostgreSQL) Close() {
	p.db.Close()
}

func (p *PostgreSQL) GetDB() *sqlx.DB {
	return p.db
}

package db

import (
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
func NewDatabase() (*PostgreSQL, error) {
	connectionStr := "postgres://root:password@localhost:5432/lendmyspace?sslmode=disable"
	db, err := sqlx.Open("postgres", connectionStr)
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

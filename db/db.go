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

type PostgreSQLsqlx struct {
	db *sqlx.DB
}

// NewDatabase make new database connection
func NewDatabase() (*PostgreSQLsqlx, error) {
	connectionStr := "postgres://root:password@dplatform_postgres:5432/dplatform?sslmode=disable"
	db, err := sqlx.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(maxConnLife)

	return &PostgreSQLsqlx{
		db: db,
	}, nil
}

func (p *PostgreSQLsqlx) Close() {
	p.db.Close()
}

func (p *PostgreSQLsqlx) GetDB() *sqlx.DB {
	return p.db
}

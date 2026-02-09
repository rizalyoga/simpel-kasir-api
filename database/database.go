package database

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	config, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db := stdlib.OpenDB(*config)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}

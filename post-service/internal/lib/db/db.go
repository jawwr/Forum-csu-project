package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type Db struct {
	PgConn *pgxpool.Pool
}

func New(ctx context.Context) *Db {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		"disable",
	)

	conn, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		log.Fatal(err)
	}

	return &Db{PgConn: conn}
}

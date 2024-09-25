package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectPostgres(ctx context.Context) *pgxpool.Pool {
	dsn := "postgres://user:password@localhost:5432/support_db"
	db, err := pgxpool.Connect(ctx, dsn)

	if err != nil {
		log.Fatalf("Неу удалось подключиться к бд")
		return nil
	}
	return db
}

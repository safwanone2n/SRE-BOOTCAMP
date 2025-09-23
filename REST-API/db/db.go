package db

import (
	"context"
	"os"
	"rest-api/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConn(ctx context.Context) (*pgxpool.Pool, error) {
	pgx, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return pgx, nil
}

func GetQuerier(ctx context.Context) (sqlc.Querier, error) {
	pgxPool, err := GetConn(ctx)
	if err != nil {
		return nil, err
	}
	return sqlc.New(pgxPool), nil
}

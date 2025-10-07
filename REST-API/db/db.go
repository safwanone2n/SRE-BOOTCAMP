package db

import (
	"context"
	"rest-api/config"
	"rest-api/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConn(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	pgx, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err := pgx.Ping(ctx); err != nil {
		return nil, err
	}
	return pgx, nil
}

func GetQuerier(ctx context.Context, cfg *config.Config) (sqlc.Querier, error) {
	pgxPool, err := GetConn(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return sqlc.New(pgxPool), nil
}

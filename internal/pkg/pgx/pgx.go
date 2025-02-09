package pkg

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/relationskatie/timetotest/config"
	"go.uber.org/zap"
)

func New(cfg *config.Config, logger *zap.Logger) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	c, err := pgxpool.ParseConfig(cfg.Postgres.GetAddress())
	if err != nil {
		return nil, err
	}
	pool, err = pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return nil, err
	}
	logger.Info("created sql client")
	return pool, nil
}

package pgx

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/relationskatie/timetotest/internal/storage"
	"go.uber.org/zap"
)

var _ storage.Interface = (*Storage)(nil)

type Storage struct {
	user *user
	log  *zap.Logger
	pgx  *pgxpool.Pool
}

func New(log *zap.Logger, pgx *pgxpool.Pool) *Storage {
	user := NewUser(log, pgx)
	return &Storage{
		user: user,
		log:  log,
		pgx:  pgx,
	}

}

func (s *Storage) User() storage.UserStorage {
	return s.user
}

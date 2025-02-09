package pgx

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/relationskatie/timetotest/internal/modles"
	"github.com/relationskatie/timetotest/internal/storage"
	"go.uber.org/zap"
)

var _ storage.UserStorage = (*user)(nil)

type user struct {
	log *zap.Logger
	pgx *pgxpool.Pool
}

func NewUser(log *zap.Logger, pgx *pgxpool.Pool) *user {
	user := &user{log: log, pgx: pgx}
	user.log.Info("initialized user repository")
	return user
}

func (u *user) DeleteUserByUsername(ctx context.Context, username string) error {
	const s = `DELETE FROM users WHERE username = $1`
	_, err := u.pgx.Exec(ctx, s, username)
	if err != nil {
		u.log.Error("failed to delete user", zap.String("username", username), zap.Error(err))
		return err
	}

	return nil
}

func (u *user) GetAllUsers(ctx context.Context) ([]modles.UserDTO, error) {
	var users []modles.UserDTO
	const s = `SELECT * FROM users`
	rows, err := u.pgx.Query(ctx, s)
	if err != nil {
		u.log.Error("failed to query all users", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user modles.UserDTO
		err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Age, &user.Telephone)
		if err != nil {
			u.log.Error("failed to query all users", zap.Error(err))
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		u.log.Error("failed to query all users", zap.Error(err))
		return nil, err
	}
	return users, nil
}

func (u *user) ChangeUser(ctx context.Context, dto modles.ChangeUserDTO) error {
	const s = `UPDATE users SET name = $1, age = $2, telephone = $3 WHERE username = $4`
	_, err := u.pgx.Exec(ctx, s, dto.Name, dto.Age, dto.Telephone, dto.Username)
	if err != nil {
		u.log.Error("failed to update user", zap.String("username", dto.Username), zap.Error(err))
		return err
	}
	return nil
}

func (u *user) AddNewUser(ctx context.Context, user modles.UserDTO) error {
	const s = `SELECT COUNT(*) FROM users WHERE username = $1`
	const t = `INSERT INTO users (id, name, username, age, telephone) VALUES ($1, $2, $3, $4, $5)`
	var cnt int
	err := u.pgx.QueryRow(ctx, s, user.Username).Scan(&cnt)
	if err != nil {
		u.log.Error("failed to add new user", zap.String("username", user.Username), zap.Error(err))
		return err
	}
	if cnt > 0 {
		u.log.Info("user already exists", zap.String("username", user.Username), zap.Int("count", cnt))
		return errors.New("user already exists")
	}
	_, err = u.pgx.Exec(ctx, t, user.ID, user.Name, user.Username, user.Age, user.Telephone)
	if err != nil {
		u.log.Error("failed to add new user", zap.String("username", user.Username), zap.Error(err))
		return err
	}
	return nil
}

func (u *user) GetUserByID(ctx context.Context, ID uuid.UUID) (modles.UserDTO, error) {
	if ID == uuid.Nil {
		u.log.Error("empty user ID")
		return modles.UserDTO{}, fmt.Errorf("invalid user ID")
	}

	const s = `SELECT * FROM users WHERE id = $1`
	var res modles.UserDTO
	err := u.pgx.QueryRow(ctx, s, ID).Scan(&res.ID, &res.Name, &res.Username, &res.Age, &res.Telephone)
	if err != nil {
		if err == pgx.ErrNoRows {
			u.log.Error("user not found", zap.String("id", ID.String()))
			return modles.UserDTO{}, fmt.Errorf("user with id %s not found", ID.String())
		}
		u.log.Error("failed to query user", zap.String("id", ID.String()), zap.Error(err))
		return modles.UserDTO{}, fmt.Errorf("failed to query user with id %s: %w", ID.String(), err)
	}
	return res, nil
}

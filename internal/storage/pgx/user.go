package pgx

import (
	"context"
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
		err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Age, *&user.Telephone)
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

func (u *user) ChangeUser(ctx context.Context, username string) error {
	const s = `UPDATE users SET name = $1, age = $2, telephone = $3 WHERE username = $4`
	_, err := u.pgx.Exec(ctx, s, username, username, username, username, username)
	if err != nil {
		u.log.Error("failed to update user", zap.String("username", username), zap.Error(err))
		return err
	}
	return nil
}

func (u *user) AddNewUser(ctx context.Context, user modles.UserDTO) error {
	const s = `SELECT * FROM users WHERE username = $1`
	const t = `INSERT INTO users (id, name, username, age, telephone) VALUES ($1, $2, $3, $4, $5)`
	_, err := u.pgx.Exec(ctx, s, user.Username)
	if err != nil {
		u.log.Error("failed to add new user", zap.String("username", user.Username), zap.Error(err))
		return err
	}
	_, err = u.pgx.Exec(ctx, t, user.ID, user.Name, user.Age, user.Telephone)
	if err != nil {
		u.log.Error("failed to add new user", zap.String("username", user.Username), zap.Error(err))
		return err
	}
	return nil
}

package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/relationskatie/timetotest/internal/modles"
)

type UserStorage interface {
	DeleteUserByUsername(ctx context.Context, username string) error
	GetAllUsers(ctx context.Context) ([]modles.UserDTO, error)
	ChangeUser(ctx context.Context, dto modles.ChangeUserDTO) error
	AddNewUser(ctx context.Context, user modles.UserDTO) error
	GetUserByID(ctx context.Context, ID uuid.UUID) (modles.UserDTO, error)
}

type Interface interface {
	User() UserStorage
}

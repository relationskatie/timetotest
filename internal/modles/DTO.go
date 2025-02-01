package modles

import "github.com/google/uuid"

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	Telephone string    `json:"telephone"`
}

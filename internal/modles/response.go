package modles

import "github.com/google/uuid"

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	Telephone string    `json:"telephone"`
}

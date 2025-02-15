package modles

import "github.com/google/uuid"

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	Telephone string    `json:"telephone"`
}
type ChangeUserDTO struct {
	Name      string `json:"name"`
	Age       uint   `json:"age"`
	Telephone string `json:"telephone"`
	Username  string `json:"username"`
}

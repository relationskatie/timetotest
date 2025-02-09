package modles

type AddUserRequest struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Age       uint   `json:"age"`
	Telephone string `json:"telephone"`
}
type ChangeUserRequest struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Age       uint   `json:"age"`
	Telephone string `json:"telephone"`
}

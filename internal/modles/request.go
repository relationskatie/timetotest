package modles

type AddUserRequest struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Telephone string `json:"telephone"`
}

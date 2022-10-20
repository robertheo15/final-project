package response

type UserCreateResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

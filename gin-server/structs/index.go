package structs

type RegisterUser struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisteredUserResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type LoginUser struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

package model

type UserCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}

package models

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	
}
type LoginResponse struct {
	Status int8 `json:"status"`
	UserData UserData `json:"user_data"`
}
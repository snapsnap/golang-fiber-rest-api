package dto

type UserData struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

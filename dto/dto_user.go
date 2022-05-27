package dto

type UserDto struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Id       int    `json:"Id"`
}

type UsersDto []UserDto

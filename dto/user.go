package dto

type UserDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Password int    `json:"password"`
	Id       int    `json:"id"`
}
type UsersDto []UserDto

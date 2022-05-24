package dto

type UserReqDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Id       int    `json:"id"`
}

type UsersReqDto []UserReqDto
package dto

type UserRespDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Id       int    `json:"id"`
}

type UsersRespDto []UserRespDto
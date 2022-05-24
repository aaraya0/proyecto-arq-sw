package dto

type CartDto struct {
	Name     UsersReqDto `json:"name"`
	LastName UsersReqDto `json:"last_name"`
	Product  ProductsDto `json:"product"`
	
	Id       int    `json:"id"`
}

type CartsDto []CartDto
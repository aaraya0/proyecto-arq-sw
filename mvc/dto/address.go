package dto

type AddressDto struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Number  int    `json:"number"`
	Id      int    `json:"id"`
	CP      int    `json:"cod_postal"`
	User_Id int    `json:"user_id"`
}

type AddressesDto []AddressDto

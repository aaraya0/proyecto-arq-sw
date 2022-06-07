package dto

type CartDto struct {
	IdCart int    `json:"id_cart"`
	Id     int    `json:"id_prod"`
	Title  string `json:"title"`
	Author string `json:"author"`
	//UniversalCode string  `json:"universal_code"`
	Price float32 `json:"base_price"`
	Stock int     `json:"stock"`

	//Picture       string  //`json:"picture_url"`
}

type CartsDto []CartDto

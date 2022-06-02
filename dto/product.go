package dto

type ProductDto struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	//UniversalCode string  `json:"universal_code"`
	Price float32 `json:"base_price"`
	Stock int     `json:"stock"`

	//Picture       string  //`json:"picture_url"`
}

type ProductsDto []ProductDto

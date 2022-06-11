package dto

type DetailDto struct {
	Quantity  int     `json:"quantity"`
	Product []ProductDto `json:"product"`
	Price   float32      `json:"price"`
	Id      int          `json:"id"`
}

type DetailsDto []DetailDto

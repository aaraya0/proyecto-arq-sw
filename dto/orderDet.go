package dto

type DetailDto struct {
	Amount  float32      `json:"amount"`
	Product []ProductDto `json:"product"`
	Price   float32      `json:"price"`
	Id      int          `json:"id"`
}

type DetailsDto []DetailDto

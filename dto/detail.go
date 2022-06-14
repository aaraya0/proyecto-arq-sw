package dto

type DetailDto struct {
	Quantity     int     `json:"quantity"`
	Product_Id   int     `json:"product_id"`
	Price        float32 `json:"price"`
	Id           int     `json:"id"`
	Product_Name string  `json:"product_name"`
}

type DetailsDto []DetailDto

type OrderDetailIDto struct {
	Product_Id   int     `json:"product_id"`
	Quantity     int     `json:"quantity"`
	Price        float32 `json:"price"`
	Product_Name string  `json:"product_name"`
}
type OrderDetailsIDto []OrderDetailIDto

type OrderDetailResponseDto struct {
	Id int `json:"detail_id"`
}

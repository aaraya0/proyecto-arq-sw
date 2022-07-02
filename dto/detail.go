package dto

type DetailDto struct {
	Id         int     `json:"id"`
	UnitPrice  float32 `json:"unit_price"`
	Quantity   float32 `json:"quantity"`
	Total      float32 `json:"total"`
	Name       string  `json:"name"`
	Product_Id int     `json:"product_id"`
	Order_Id   int     `json:"order_id"`
}

type DetailsDto []DetailDto

/*
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
*/

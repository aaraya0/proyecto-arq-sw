package dto

type OrderDto struct {
	Date        string    `json:"order_date"` //puede no estar
	Total       float32   `json:"total"`
	OrderDetail DetailDto `json:"detail"`
	Id          int       `json:"id"`
}

type OrdersDto []OrderDto

type OrderDtoInsert struct {
	User_Id      int              `json:"user_id"`
	OrderDetails OrderDetailsIDto `json:"details"`
}

type OrderDtoResp struct {
	Id int `json:"order_id"`
}

package dto

type OrderDto struct {
	//OrderDate      string     `json: "order_date"`//puede no estar
	Total       float32     `json:"total"`
	OrderDetail []DetailDto `json:"detail"`
	User        []UserDto   `json:"user"`
	Id          int         `json:"id"`
}

type OrdersDto []OrderDto

type OrderDtoResp struct {
	Id int `json:"order_id"`
}

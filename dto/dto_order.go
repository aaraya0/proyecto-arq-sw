package dto

type OrderDto struct {
	OrderDate      string     `json: "order_date"`//puede no estar
	Total          float32    `json: "total"`
	OrderDetailDto DetailsDto `json: "orden"`
	UserDto        UsersReqDto   `json: "user"`
	Id             int        `json: "id"`
}

type OrdersDto [] OrderDto
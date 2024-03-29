package dto

import (
	"time"
)

type OrderDto struct {
	Id           int        `json:"order_id"`
	TotalAmount  float32    `json:"total_amount"`
	Date         time.Time  `json:"date"`
	User_Id      int        `json:"user_id"`
	OrderDetails DetailsDto `json:"orderDetail"`
}

type OrdersDto []OrderDto

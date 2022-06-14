package services

import (
	//falta client
	"time"

	orderCliente "github.com/aaraya0/arq-software/Integrador1/clients/order"
	productCliente "github.com/aaraya0/arq-software/Integrador1/clients/product"
	"github.com/aaraya0/arq-software/Integrador1/dto"
	"github.com/aaraya0/arq-software/Integrador1/model"
	e "github.com/aaraya0/arq-software/Integrador1/utils/errors"
)

type orderService struct{}

type orderServiceInterface interface {
	AddOrder(orderDto dto.OrderDtoInsert) (dto.OrderDtoResp, e.ApiError)
	GetOrdersByUId(id int) (dto.OrdersDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}
func (s *orderService) GetOrdersByUId(id int) (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderCliente.GetOrdersByUId(id)
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto
		orderDto.Id = order.Id
		orderDto.Date = order.Date
		orderDto.Total = order.Total
		orderDto.OrderDetail, _ = DetailService.GetOrderDetailById(order.Id)
		ordersDto = append(ordersDto, orderDto)
	}
	return ordersDto, nil
}

func (s *orderService) AddOrder(orderDtoInsert dto.OrderDtoInsert) (dto.OrderDtoResp, e.ApiError) {

	var order model.Order
	var total float32
	var orderDtoResp dto.OrderDtoResp
	total = 0
	order.User_id = orderDtoInsert.User_Id
	order.Date = time.Now().Format("2022.03.04 17:06:04")
	for i := 0; i < len(orderDtoInsert.OrderDetails); i++ {
		var product model.Product
		detail := orderDtoInsert.OrderDetails[i]
		product = productCliente.GetProductById(detail.Product_Id)
		if product.Stock < detail.Quantity {
			orderDtoResp.Id = 0
			return orderDtoResp, e.NewConflictApiError("No hay stock de:" + product.Title)
		}

		total += (detail.Price * float32(detail.Quantity))

	}

	for i := 0; i < len(orderDtoInsert.OrderDetails); i++ {
		detail := orderDtoInsert.OrderDetails[i]
		productCliente.SubsStock(detail.Product_Id, detail.Quantity)
	}

	order.Total = total

	order = orderCliente.AddOrder(order)

	orderDtoResp.Id = order.Id

	return orderDtoResp, nil
}

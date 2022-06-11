package services

import (
	//falta client
	ordCliente "github.com/aaraya0/arq-software/Integrador1/clients/order"
	"github.com/aaraya0/arq-software/Integrador1/dto"
	"github.com/aaraya0/arq-software/Integrador1/model"
	e "github.com/aaraya0/arq-software/Integrador1/utils/errors"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrderById(id int) (dto.OrderDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	//InsertProduct
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}
func (p *orderService) GetOrderById(id int) (dto.OrderDto, e.ApiError) {
	var order model.Order = ordCliente.GetOrderById(id)
	var orderDto dto.OrderDto
	if order.Id == 0 {
		return orderDto, e.NewBadRequestApiError("order not found")
	}
	
	orderDto.Id = order.Id
	//orderDto.User = order.User
	orderDto.Total = order.Total
	//orderDto.OrderDetail = order.OrderDetail	
	return orderDto, nil
}
func (p *orderService) GetOrders() (dto.OrdersDto, e.ApiError) {
	var orders model.Orders = ordCliente.GetOrders()
	var ordersDto dto.OrdersDto
	for _, order := range orders {
		var orderDto dto.OrderDto
		orderDto.Id = order.Id
	    //orderDto.User = order.User
	    orderDto.Total = order.Total
	    //orderDto.OrderDetail = order.OrderDetail
	
		
		ordersDto = append(ordersDto, orderDto)
	}
	return ordersDto, nil

}
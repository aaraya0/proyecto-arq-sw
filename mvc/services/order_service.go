package services

import (
	"time"

	orderDetailCliente "mvc/clients/detail"
	orderCliente "mvc/clients/order"
	productCliente "mvc/clients/product"
	"mvc/dto"
	"mvc/model"
	e "mvc/utils/errors"
)

type orderService struct{}

type orderServiceInterface interface {
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
	GetOrdersByUserId(token int) (dto.OrdersDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError) {

	var order model.Order

	var date = time.Now()

	order.Date = date
	order.User_Id = orderDto.User_Id

	order = orderCliente.InsertOrder(order)

	orderDto.Id = order.Id

	var orderDetails model.Details
	var montofinal float32
	for _, orderDetailDto := range orderDto.OrderDetails {
		var orderDetail model.Detail

		orderDetail.Product_Id = orderDetailDto.Product_Id

		var product model.Product = productCliente.GetProductById(orderDetail.Product_Id)
		orderDetail.Name = product.Title
		orderDetail.UnitPrice = product.Price

		orderDetail.Quantity = orderDetailDto.Quantity
		orderDetail.Total = orderDetail.UnitPrice * orderDetail.Quantity

		orderDetail.Order_Id = orderDto.Id
		for i := 0; i < len(orderDto.OrderDetails); i++ {
			detail := orderDto.OrderDetails[i]
			product := productCliente.GetProductById(detail.Product_Id)
			if float32(product.Stock) < detail.Quantity {
				orderDto.Id = 0
				return orderDto, e.NewConflictApiError("Not enough stock on product: " + product.Title)
			}

		}
		montofinal += orderDetail.Total
		orderDetails = append(orderDetails, orderDetail)
	}

	order.TotalAmount = orderCliente.UpdateMontoFinal(montofinal, order.Id)

	orderDetails = orderDetailCliente.InsertOrderDetails(orderDetails)

	var orderResponseDto dto.OrderDto

	orderResponseDto.Date = order.Date
	orderResponseDto.Id = order.Id
	orderResponseDto.User_Id = order.User_Id
	orderResponseDto.TotalAmount = order.TotalAmount

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.DetailDto

		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Quantity = orderDetail.Quantity
		orderDetailDto.Product_Id = orderDetail.Product_Id
		orderDetailDto.UnitPrice = orderDetail.UnitPrice
		orderDetailDto.Total = orderDetail.Total
		orderDetailDto.Order_Id = orderDetail.Order_Id
		orderDetailDto.Name = orderDetail.Name

		orderResponseDto.OrderDetails = append(orderResponseDto.OrderDetails, orderDetailDto)
	}

	return orderResponseDto, nil
}

//Buscar orden por IDuser

func (s *orderService) GetOrdersByUserId(token int) (dto.OrdersDto, e.ApiError) {

	var IdUserX int = token
	var orders model.Orders = orderCliente.GetOrdersByUserId(IdUserX)
	var ordersDto dto.OrdersDto

	if len(orders) == 0 {
		return ordersDto, e.NewBadRequestApiError("order not found")
	}

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id = order.Id
		orderDto.Date = order.Date
		orderDto.TotalAmount = order.TotalAmount
		orderDto.User_Id = order.User_Id
		var orderDetails model.Details = orderDetailCliente.GetOrderDetailByOrderId(order.Id)
		for _, orderDetail := range orderDetails {
			var orderDetailDto dto.DetailDto

			orderDetailDto.Id = orderDetail.Id
			orderDetailDto.Quantity = orderDetail.Quantity
			orderDetailDto.Product_Id = orderDetail.Product_Id
			orderDetailDto.UnitPrice = orderDetail.UnitPrice
			orderDetailDto.Total = orderDetail.Total
			orderDetailDto.Order_Id = orderDetail.Order_Id
			orderDetailDto.Name = orderDetail.Name

			orderDto.OrderDetails = append(orderDto.OrderDetails, orderDetailDto)
		}
		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

package services

import (
	orderDetailCliente "mvc/clients/detail"
	"mvc/dto"
	"mvc/model"
	e "mvc/utils/errors"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	GetOrderDetailByOrderId(orderId int) (dto.DetailsDto, e.ApiError)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

func (s *orderDetailService) GetOrderDetailByOrderId(orderId int) (dto.DetailsDto, e.ApiError) {

	var orderDetails model.Details = orderDetailCliente.GetOrderDetailByOrderId(orderId)
	var orderDetailsResDto dto.DetailsDto

	if len(orderDetails) == 0 {
		return orderDetailsResDto, e.NewBadRequestApiError("Detail not found")
	}
	for _, orderDetailRes := range orderDetails {
		var orderDetailResDto dto.DetailDto
		orderDetailResDto.Id = orderDetailRes.Id
		orderDetailResDto.Name = orderDetailRes.Name
		orderDetailResDto.Quantity = orderDetailRes.Quantity
		orderDetailResDto.UnitPrice = orderDetailRes.UnitPrice
		orderDetailResDto.Total = orderDetailRes.Total
		orderDetailResDto.Product_Id = orderDetailRes.Product_Id
		orderDetailResDto.Order_Id = orderDetailRes.Order_Id
		orderDetailsResDto = append(orderDetailsResDto, orderDetailResDto)
	}
	return orderDetailsResDto, nil
}

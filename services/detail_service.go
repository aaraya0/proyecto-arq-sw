package services

import (
	//falta client
	detCliente "github.com/aaraya0/arq-software/Integrador1/clients/detail"
	"github.com/aaraya0/arq-software/Integrador1/dto"
	"github.com/aaraya0/arq-software/Integrador1/model"
	e "github.com/aaraya0/arq-software/Integrador1/utils/errors"
)

type detailService struct{}

type detailServiceInterface interface {
	AddOrderDetail(OrderDetailIDto dto.OrderDetailIDto, orderId int) (dto.OrderDetailResponseDto, e.ApiError)
	GetOrderDetailById(id int) (dto.DetailDto, e.ApiError)
	GetDetailsByOId(id int) (dto.DetailsDto, e.ApiError)
}

var (
	DetailService detailServiceInterface
)

func init() {
	DetailService = &detailService{}
}
func (s *detailService) GetOrderDetailById(id int) (dto.DetailDto, e.ApiError) {

	var detail model.Detail = detCliente.GetDetailById(id)
	var detailDto dto.DetailDto

	detailDto.Id = detail.Id
	detailDto.Product_Id = detail.Product_id
	detailDto.Quantity = detail.Quantity
	detailDto.Price = detail.Price
	detailDto.Product_Name = detail.Product_name

	return detailDto, nil
}

func (s *detailService) GetDetailsByOId(id int) (dto.DetailsDto, e.ApiError) {
	var details model.Details = detCliente.GetDetailsByOId(id)
	var detailsDto dto.DetailsDto

	for _, orderDetail := range details {
		var detailDto dto.DetailDto
		detailDto.Id = orderDetail.Id
		detailDto.Product_Id = orderDetail.Product_id
		detailDto.Quantity = orderDetail.Quantity
		detailDto.Price = orderDetail.Price
		detailDto.Product_Name = orderDetail.Product_name

		detailsDto = append(detailsDto, detailDto)
	}

	return detailsDto, nil
}

func (s *detailService) AddOrderDetail(detailDto dto.OrderDetailIDto, orderId int) (dto.OrderDetailResponseDto, e.ApiError) {

	var detail model.Detail
	detail.Order_id = orderId
	detail.Product_id = detailDto.Product_Id
	detail.Quantity = detailDto.Quantity
	detail.Product_name = detailDto.Product_Name
	detail.Price = detailDto.Price

	detail = detCliente.AddOrderDetail(detail)

	var orderDetailResponseDto dto.OrderDetailResponseDto
	orderDetailResponseDto.Id = detail.Id

	return orderDetailResponseDto, nil
}

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
	GetDetailById(id int) (dto.DetailDto, e.ApiError)
	GetDetails() (dto.DetailsDto, e.ApiError)
	//InsertProduct
}

var (
	DetailService detailServiceInterface
)

func init() {
	DetailService = &detailService{}
}
func (p *detailService) GetDetailById(id int) (dto.DetailDto, e.ApiError) {
	var detail model.Detail = detCliente.GetDetailById(id)
	var detailDto dto.DetailDto
	if detail.Id == 0 {
		return detailDto, e.NewBadRequestApiError("detail not found")
	}
	
	detailDto.Id = detail.Id
	detailDto.Quantity = detail.Quantity
	//detailDto.Product = detail.Product
	detailDto.Price = detail.Price
	
	return detailDto, nil
}
func (p *detailService) GetDetails() (dto.DetailsDto, e.ApiError) {
	var details model.Details = detCliente.GetDetails()
	var detailsDto dto.DetailsDto
	for _, detail := range details {
		var detailDto dto.DetailDto
		detailDto.Id = detail.Id
	detailDto.Quantity = detail.Quantity
	//detailDto.Product = detail.Product
	detailDto.Price = detail.Price
	
		
	detailsDto = append(detailsDto, detailDto)
	}
	return detailsDto, nil

}
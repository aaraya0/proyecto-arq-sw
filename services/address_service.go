package services

import (
	addressCliente "Integrador1/clients/user"
	"Integrador1/dto"
	"Integrador1/model"
	e "Integrador1/utils/errors"
)

type addressService struct{}

type addressServiceInterface interface {
	GetAddressById(id int) (dto.AddressDto, e.ApiError)
	//GetAddresses() (dto.AddressesDto, e.ApiError)
	InsertAddress(addressDto dto.AddressDto) (dto.AddressDto, e.ApiError)
}

var (
	AddressService addressServiceInterface
)

func init() {
	AddressService = &addressService{}
}

func (s *addressService) GetAddressById(id int) (dto.AddressDto, e.ApiError) {

	var address model.Address = addressCliente.GetAdressById(id)
	var addressDto dto.AddressDto

	if address.Id == 0 {
		return addressDto, e.NewBadRequestApiError("address not found")
	}
	addressDto.StName = address.StName
	addressDto.StNumber = address.StNumber

	addressDto.Id = address.Id
	return addressDto, nil
}

/*func (s *addressService) GetAddresses() (dto.AddressesDto, e.ApiError) {

	var addresses model.Addresses = addressCliente.GetAddresses()
	var addressesDto dto.AddressesDto

	for _, address := range addresses {
		var addressDto dto.AddressDto
		addressDto.StName = address.StName
		addressDto.StNumber = address.StNumber

		addressDto.Id = address.Id
		addressesDto = append(addressesDto, addressDto)
	}

	return addressesDto, nil
}
*/
func (s *addressService) InsertAddress(addressDto dto.AddressDto) (dto.AddressDto, e.ApiError) {

	var address model.Address

	addressDto.StName = address.StName
	addressDto.StNumber = address.StNumber

	address = addressCliente.InsertAddress(address)

	addressDto.Id = address.Id

	return addressDto, nil
}

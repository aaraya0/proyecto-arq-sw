package services

import (
	addrCliente "github.com/aaraya0/arq-software/Integrador1/mvc/clients/address"

	"github.com/aaraya0/arq-software/Integrador1/mvc/dto"
	"github.com/aaraya0/arq-software/Integrador1/mvc/model"
	e "github.com/aaraya0/arq-software/Integrador1/mvc/utils/errors"
	//"github.com/dgrijalva/jwt-go"
)

type addressService struct{}

type addressServiceInterface interface {
	InsertAddress(addrDto dto.AddressDto) (dto.AddressDto, e.ApiError)
	GetAddrByUserId(token int) (dto.AddressesDto, e.ApiError) //token cambiado de string a int temporalmente. revisar
}

var (
	AddressService addressServiceInterface
)

func init() {
	AddressService = &addressService{}
}

func (s *addressService) InsertAddress(addrDto dto.AddressDto) (dto.AddressDto, e.ApiError) {

	var addr model.Address

	addr.User_Id = addrDto.User_Id
	addr.CP = addrDto.CP
	addr.City = addrDto.City
	addr.Country = addrDto.Country
	addr.Street = addrDto.Street
	addr.Number = addrDto.Number

	addr = addrCliente.InsertAddress(addr)

	addrDto.Id = addr.Id

	var addrResponseDto dto.AddressDto

	addrResponseDto.Id = addr.Id
	addrResponseDto.User_Id = addr.User_Id
	addrResponseDto.CP = addr.CP
	addrResponseDto.City = addr.City
	addrResponseDto.Country = addr.Country
	addrResponseDto.Number = addr.Number
	addrResponseDto.Street = addr.Street

	return addrResponseDto, nil
}

//Buscar orden por IDuser

func (s *addressService) GetAddrByUserId(token int) (dto.AddressesDto, e.ApiError) { //token cambiado de string a int temporalmente (login sin token)
	//var userId float64
	/*tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, e.NewUnauthorizedApiError("Unauthorized")
		}
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}

	if !tkn.Valid {
		return nil, e.NewUnauthorizedApiError("Unauthorized")

	}
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {

		userId = (claims["user_id"].(float64))

	} else {
		return nil, e.NewUnauthorizedApiError("Unauthorized")
	}
	*/
	//var IdUserX int = int(userId)   //temporalmente comentado porq el login es sin token. revisar
	var IdUserX int = token
	var addrs model.Addresses = addrCliente.GetAddrByUserId(IdUserX)
	var addrsDto dto.AddressesDto

	if len(addrs) == 0 {
		return addrsDto, e.NewBadRequestApiError("address not found")
	}

	for _, addr := range addrs {
		var addrDto dto.AddressDto
		addrDto.User_Id = addr.User_Id
		addrDto.Id = addr.Id
		addrDto.CP = addr.CP
		addrDto.City = addr.City
		addrDto.Country = addr.Country
		addrDto.Street = addr.Street
		addrDto.Number = addr.Number

		addrsDto = append(addrsDto, addrDto)
	}

	return addrsDto, nil
}

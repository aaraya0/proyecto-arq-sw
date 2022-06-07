package services

import (
	cartCliente "github.com/aaraya0/arq-software/Integrador1/clients/cart"
	"github.com/aaraya0/arq-software/Integrador1/dto"
	"github.com/aaraya0/arq-software/Integrador1/model"
	e "github.com/aaraya0/arq-software/Integrador1/utils/errors"
)

type cartService struct{}

type cartServiceInterface interface {
	AddProduct(cartDto dto.CartDto) (dto.CartDto, e.ApiError)
	GetCart() (dto.CartsDto, e.ApiError)
	//InsertProduct
}

var (
	CartService cartServiceInterface
)

func init() {
	CartService = &cartService{}
}
func (s *cartService) AddProduct(cartDto dto.CartDto) (dto.CartDto, e.ApiError) {

	var cart model.Cart
	cartDto.Title = cart.Title
	cartDto.Author = cart.Author
	cartDto.Price = cart.Price
	//cartDto.Stock = cart.Stock
	//productDto.UniversalCode = product.UniversalCode
	cartDto.Id = cart.Id
	cartDto.IdCart = cart.IdCart

	cart = cartCliente.AddProduct(cart)

	return cartDto, nil
}
func (p *cartService) GetCart() (dto.CartsDto, e.ApiError) {
	var carts model.Carts = cartCliente.GetCart()
	var cartsDto dto.CartsDto
	for _, cart := range carts {
		var cartDto dto.CartDto
		cartDto.Title = cart.Title
		cartDto.Author = cart.Author
		cartDto.Price = cart.Price
		//cartDto.Stock = cart.Stock
		//productDto.UniversalCode = product.UniversalCode
		cartDto.Id = cart.Id
		cartDto.IdCart = cart.IdCart
		cartsDto = append(cartsDto, cartDto)
	}
	return cartsDto, nil

}

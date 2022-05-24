package services

import (
	productCliente "Integrador1/clients/product"
	"Integrador1/dto"
	"Integrador1/model"
	e "Integrador1/utils/errors"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	InsertProduct(productDto dto.ProductDto) (dto.ProductDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id)
	var productDto dto.ProductDto

	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}
	productDto.product_name = product.Product_Name

	productDto.Id = product.Id
	return productDto, nil
}

func (s *productService) GetProducts() (dto.productsDto, e.ApiError) {

	var products model.products = productCliente.GetProducts()
	var productsDto dto.productsDto

	for _, product := range products {
		var productDto dto.productDto
		productDto.product_name = product.Product_Name

		productDto.Id = product.Id
		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

func (s *productService) InsertProduct(productDto dto.productDto) (dto.productDto, e.ApiError) {

	var product model.product

	product.Name = productDto.Name
	product.LastName = productDto.LastName
	product.productName = productDto.productName
	product.Password = productDto.Password

	product = productCliente.Insertproduct(product)

	productDto.Id = product.Id

	return productDto, nil
}

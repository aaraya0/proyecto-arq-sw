package services

import (
	prodCliente "mvc/clients/product"
	"mvc/dto"
	"mvc/model"
	e "mvc/utils/errors"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	GetProductsByCategoryId(id int) (dto.ProductsDto, e.ApiError)
	//InsertProduct
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}
func (p *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {
	var product model.Product = prodCliente.GetProductById(id)
	var productDto dto.ProductDto
	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}
	productDto.Title = product.Title
	productDto.Author = product.Author
	productDto.Price = product.Price
	productDto.Stock = product.Stock
	productDto.Image = product.Image
	productDto.Category_id = product.Category_id
	productDto.Id = product.Id
	return productDto, nil
}
func (p *productService) GetProducts() (dto.ProductsDto, e.ApiError) {
	var products model.Products = prodCliente.GetProducts()
	var productsDto dto.ProductsDto
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Title = product.Title
		productDto.Author = product.Author
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.Id = product.Id
		productDto.Image = product.Image
		productDto.Category_id = product.Category_id
		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil

}
func (s *productService) GetProductsByCategoryId(id int) (dto.ProductsDto, e.ApiError) {

	var products model.Products = prodCliente.GetProductsByCategoryId(id)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id = product.Id
		//productDto.Category, _ = CategoryService.GetCategoryById(id)
		productDto.Category_id = product.Category_id
		productDto.Title = product.Title
		productDto.Author = product.Author
		productDto.Price = product.Price
		productDto.Image = product.Image
		productDto.Stock = product.Stock

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

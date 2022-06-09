package services

import (
	categCliente "github.com/aaraya0/arq-software/Integrador1/clients/category"
	"github.com/aaraya0/arq-software/Integrador1/dto"
	"github.com/aaraya0/arq-software/Integrador1/model"
	e "github.com/aaraya0/arq-software/Integrador1/utils/errors"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategoryById(id int) (dto.CategoryDto, e.ApiError)
	GetCategories() (dto.CategoriesDto, e.ApiError)
	//InsertProduct

}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}
func (p *categoryService) GetCategoryById(id int) (dto.CategoryDto, e.ApiError) {
	var category model.Category = categCliente.GetCategoryById(id)
	var categoryDto dto.CategoryDto
	if category.Id == 0 {
		return categoryDto, e.NewBadRequestApiError("category not found")
	}
	categoryDto.Id = category.Id
	categoryDto.Name = category.Name

	//productDto.UniversalCode = product.UniversalCode

	return categoryDto, nil
}
func (p *categoryService) GetCategories() (dto.CategoriesDto, e.ApiError) {
	var categories model.Categories = categCliente.GetCategories()
	var categoriesDto dto.CategoriesDto
	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Id = category.Id
		categoryDto.Name = category.Name

		//	productDto.UniversalCode = product.UniversalCode

		categoriesDto = append(categoriesDto, categoryDto)
	}
	return categoriesDto, nil

}

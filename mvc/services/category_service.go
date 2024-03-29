package services

import (
	categCliente "mvc/clients/category"
	"mvc/dto"
	"mvc/model"
	e "mvc/utils/errors"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategoryById(id int) (dto.CategoryDto, e.ApiError)
	GetCategories() (dto.CategoriesDto, e.ApiError)
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

	return categoryDto, nil
}
func (p *categoryService) GetCategories() (dto.CategoriesDto, e.ApiError) {
	var categories model.Categories = categCliente.GetCategories()
	var categoriesDto dto.CategoriesDto
	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Id = category.Id
		categoryDto.Name = category.Name

		categoriesDto = append(categoriesDto, categoryDto)
	}
	return categoriesDto, nil

}

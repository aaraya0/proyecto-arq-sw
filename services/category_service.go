import (
	categoryCliente "Integrador1/clients/category"
	"Integrador1/dto"
	"Integrador1/model"
	e "Integrador1/utils/errors"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategoryById(id int) (dto.CategoryDto, e.ApiError)
	GetCategories() (dto.CategoriesDto, e.ApiError)
	InsertCategory(categoryDto dto.CategoryDto) (dto.CategoryDto, e.ApiError)
}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}

func (s *categoryService) GetCategoryById(id int) (dto.CategoryDto, e.ApiError) {

	var category model.Category = categoryCliente.GetCategoryById(id)
	var categoryDto dto.CategoryDto

	if category.Id == 0 {
		return categoryDto, e.NewBadRequestApiError("category not found")
	}
	categoryDto.Cat_name = category.Cat_name

	categoryDto.Id = category.Id
	return categoryDto, nil
}

func (s *categoryService) GetCategories() (dto.CategoriesDto, e.ApiError) {

	var categories model.Categories = categoryCliente.GetCategories()
	var categoriesDto dto.CategoriesDto

	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Cat_name = category.Cat_name

		categoryDto.Id = category.Id
		categoriesDto = append(categoriesDto, categoryDto)
	}

	return categoriesDto, nil
}

func (s *categoryService) InsertCategory(categoryDto dto.CategoryDto) (dto.CategoryDto, e.ApiError) {

	var category model.Category

	categoryDto.Cat_name = category.Cat_name

	category = categoryCliente.InsertCategory(category)

	categoryDto.Id = category.Id

	return categoryDto, nil
}

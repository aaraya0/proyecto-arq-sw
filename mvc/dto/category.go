package dto

type CategoryDto struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	//Products []ProductsDto
}

type CategoriesDto []CategoryDto

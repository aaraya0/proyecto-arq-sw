package dto

type CategoryDto struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
}

type CategoriesDto []CategoryDto
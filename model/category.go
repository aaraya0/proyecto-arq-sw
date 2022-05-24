package model

type Category struct {
	Id       int    `gorm:"primaryKey"`
	Cat_name   string `gorm:"type:varchar(350);not null"`
	Prods Products
}

type Categories []Category

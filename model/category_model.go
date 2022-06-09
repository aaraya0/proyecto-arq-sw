package model

type Category struct {
	Id   int    `gorm:"type:int(11)"`
	Name string `gorm:"type:varchar(350);not null"`
	//Products []Products

}

type Categories []Category

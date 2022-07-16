package model

type Address struct {
	Id      int    `gorm:"type:int"`
	Street  string `gorm:"type:varchar(550)"`
	Number  int    `gorm:"type:int"`
	Country string `gorm:"type:varchar(550)"`
	City    string `gorm:"type:varchar(550)"`
	CP      int    `gorm:"type:int"`
	User_Id int    `gorm:"type:int"`
}

type Addresses []Address

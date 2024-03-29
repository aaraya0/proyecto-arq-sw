package model

type User struct {
	Id       int    `gorm:"type:int(11);"`
	Name     string `gorm:"type:varchar(350);"`
	LastName string `gorm:"type:varchar(250);"`
	UserName string `gorm:"type:varchar(150);"`
	Password string `gorm:"type:varchar(150);"`
}

type Users []User

package user

import (
	model "mvc/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetUserByUname(uname string) model.User {
	var user model.User
	Db.Where("UserName=?", uname).First(&user)

	log.Debug("User:", user)
	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)
	log.Debug("Users: ", users)
	return users
}

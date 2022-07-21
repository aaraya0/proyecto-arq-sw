package address

import (
	model "mvc/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InsertAddress(address model.Address) model.Address {
	result := Db.Create(&address)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Address submited: ", address.Id)
	return address
}

func GetAddrByUserId(userId int) model.Addresses {
	var addrs model.Addresses

	log.Debug("userId: ", userId)
	Db.Where("user_id= ?", userId).Find(&addrs)
	log.Debug("Address: ", addrs)

	return addrs
}

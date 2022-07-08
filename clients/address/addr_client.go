package address

import (
	model "github.com/aaraya0/arq-software/Integrador1/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InsertAddress(address model.Address) model.Address {
	result := Db.Create(&address)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Address submited: ", address.Id, address.CP)
	return address
}

func GetAddrByUserId(userId int) model.Addresses {
	var addrs model.Addresses

	log.Debug("userId: ", userId)
	Db.Where("user_id= ?", userId).Find(&addrs)
	log.Debug("Address: ", addrs)

	return addrs
}
package detail

import (
	model "github.com/aaraya0/arq-software/Integrador1/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetDetailById(id int) model.Detail {
	var detail model.Detail
	Db.Where("id=?", id).First(&detail)

	log.Debug("Detail:", detail)
	return detail
}

func GetDetails() model.Details {
	var details model.Details
	Db.Find(&details)
	log.Debug("Details: ", details)
	return details

}


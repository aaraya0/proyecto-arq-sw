package category

import (
	model "github.com/aaraya0/arq-software/Integrador1/mvc/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetCategoryById(id int) model.Category {
	var category model.Category
	Db.Where("id=?", id).First(&category)

	log.Debug("Category:", category)
	return category
}

func GetCategories() model.Categories {
	var categories model.Categories
	Db.Find(&categories)
	log.Debug("Categories: ", categories)
	return categories

}

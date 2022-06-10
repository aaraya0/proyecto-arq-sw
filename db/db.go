package db

import (
	//"os"

	catCliente "github.com/aaraya0/arq-software/Integrador1/clients/category"
	prodCliente "github.com/aaraya0/arq-software/Integrador1/clients/product"
	userCliente "github.com/aaraya0/arq-software/Integrador1/clients/user"
	model "github.com/aaraya0/arq-software/Integrador1/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	dsn := "root:@tcp(127.0.0.1:3306)/sistema?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Info("Connection fail")
		log.Fatal(err)
	} else {
		log.Info("Connection success")
	}
	prodCliente.Db = db
	userCliente.Db = db
	catCliente.Db = db
}
func StartDbEngine() {
	db.AutoMigrate(&model.Product{}, &model.User{}, &model.Category{})
	log.Info("Finishing migration database tables")
}

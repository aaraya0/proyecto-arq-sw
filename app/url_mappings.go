package app

import (
	productController "github.com/aaraya0/arq-software/Integrador1/controllers"
	//userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)

	// Users Mapping
	/*	router.GET("/user/:id", userController.GetUserById)
		router.GET("/user", userController.GetUsers)
		router.POST("/user", userController.UserInsert)*/

	log.Info("Finishing mappings configurations")

}

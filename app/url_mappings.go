package app

import (
	categoryController "github.com/aaraya0/arq-software/Integrador1/controllers/category"
	productController "github.com/aaraya0/arq-software/Integrador1/controllers/product"
	userController "github.com/aaraya0/arq-software/Integrador1/controllers/user"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)
	router.GET("/products/:category_id", productController.GetProductsByCategoryId)
	// Users Mapping
	router.GET("/user/:uname", userController.GetUserByUname)
	router.GET("/user", userController.GetUsers)
	//	router.POST("/user", userController.UserInsert)
	//router.POST("/cart", cartController.AddProduct)

	//router.GET("/cart", cartController.GetCart)
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/category", categoryController.GetCategories)

	log.Info("Finishing mappings configurations")

}

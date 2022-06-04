package controllers

import (
	"net/http"
	"strconv"

	"github.com/aaraya0/arq-software/Integrador1/dto"
	service "github.com/aaraya0/arq-software/Integrador1/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Debug("Product id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var productDto dto.ProductDto
	productDto, err := service.ProductService.GetProductById(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)

}

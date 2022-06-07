package controllers

import (
	"net/http"

	"github.com/aaraya0/arq-software/Integrador1/dto"
	service "github.com/aaraya0/arq-software/Integrador1/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AddProduct(c *gin.Context) {
	var cartDto dto.CartDto
	err := c.BindJSON(&cartDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cartDto, er := service.CartService.AddProduct(cartDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, cartDto)
}

func GetCart(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var cartsDto dto.CartsDto
	cartsDto, err := service.CartService.GetCart()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, cartsDto)

}

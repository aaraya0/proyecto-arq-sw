package controllers

import (
	"net/http"
	"strconv"

	"github.com/aaraya0/arq-software/Integrador1/dto"
	service "github.com/aaraya0/arq-software/Integrador1/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderById(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Debug("Order id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var orderDto dto.OrderDto
	orderDto, err := service.OrderService.GetOrderById(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDto)
}

func GetOrders(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var ordersDto dto.OrdersDto
	ordersDto, err := service.OrderService.GetOrders()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersDto)

}
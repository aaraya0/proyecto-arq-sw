package controllers

import (
	"net/http"
	"strconv"

	"mvc/dto"
	"mvc/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderDetailByOrderId(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Debug("Detail id to load: " + c.Param("orderId"))

	orderId, _ := strconv.Atoi(c.Param("orderId"))
	var orderDetailResDto dto.DetailsDto

	orderDetailResDto, err := services.OrderDetailService.GetOrderDetailByOrderId(orderId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDetailResDto)
}

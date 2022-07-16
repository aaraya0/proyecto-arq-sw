package controllers

import (
	"net/http"
	"strconv"

	"github.com/aaraya0/arq-software/Integrador1/mvc/dto"
	service "github.com/aaraya0/arq-software/Integrador1/mvc/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AddressInsert(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "POST")
	var addrDto dto.AddressDto
	err := c.BindJSON(&addrDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	addrDto, er := service.AddressService.InsertAddress(addrDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, addrDto)
}

func GetAddrByUserId(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var addrsDto dto.AddressesDto
	token, _ := strconv.Atoi(c.Param("userId"))
	addrsDto, err := service.AddressService.GetAddrByUserId(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, addrsDto)
}

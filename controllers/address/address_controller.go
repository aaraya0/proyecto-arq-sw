package addressController

import (
	"Integrador1/dto"
	service "Integrador1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAdressById(c *gin.Context) {
	log.Debug("Address id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var addressDto dto.AddressDto

	addressDto, err := service.AddressService.GetAddressById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, addressDto)
}

/*func GetAddresses(c *gin.Context) {
	var addressesDto dto.AddressesDto
	addressesDto, err := service.AddressService.GetAddresses()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, addressesDto)
}*/

func UserInsert(c *gin.Context) {
	var addressDto dto.AddressDto
	err := c.BindJSON(&addressDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	adressDto, er := service.AddressService.InsertAdress(addressDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, addressDto)
}

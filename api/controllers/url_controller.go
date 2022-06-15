package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"tiny-url/constants"
	"tiny-url/lib"
	"tiny-url/models"
	"tiny-url/services"
)

type URLController struct {
	service services.URLService
	logger  lib.Logger
}

// NewURLController creates new url controller
func NewURLController(service services.URLService, logger lib.Logger) URLController {
	return URLController{
		service: service,
		logger:  logger,
	}
}

// GetOneUrl gets one url
func (u URLController) GetOneUrl(c *gin.Context) {
	paramID := c.Param("hash")

	if paramID == "" {
		u.logger.Error("hash is null")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "hash is null",
		})
		return
	}
	user, err := u.service.GetOneUrl(paramID)

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

// GetUrl gets all the url
func (u URLController) GetUrl(c *gin.Context) {
	users, err := u.service.GetAllUrl()

	if err != nil {
		u.logger.Error(err)
	}

	c.JSON(200, gin.H{"data": users})
}

// SaveUrl saves the url
func (u URLController) SaveUrl(c *gin.Context) {
	url := models.URL{}
	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&url); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.WithTrx(trxHandle).CreateUrl(&url); err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "url created", "hash": url.Hash})
}

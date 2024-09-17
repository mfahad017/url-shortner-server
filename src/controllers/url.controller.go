package controllers

import (
	"fmt"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/validation"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var domainPrefix = "/v1/url/"

type URLController struct{}

func (URLController) ListURLsByUser(c *gin.Context) {

}

func (URLController) GetURL(c *gin.Context) {

	index := c.Request.URL.Path[len(domainPrefix):]

	fmt.Println(c.Request.URL.Path, domainPrefix, index)

	url := models.URL{}

	err := url.GetByShortURL(index)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "URL not found",
		})
		return

	}

	oURL := "https://" + url.OriginalURL

	c.Redirect(http.StatusMovedPermanently, oURL)

}

func (URLController) CreateURL(c *gin.Context) {

	input, err := validation.Validate[validation.CreateURLInput](c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	suffix := input.Suffix

	if suffix == "" {
		suffix = "random-suffix-" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	shortUrl := config.GetConfig().DOMAIN_PREFIX + domainPrefix + suffix

	userId := 1

	url := models.URL{
		OriginalURL: input.OriginalURL,
		ShortURL:    suffix,
		UserID:      uint(userId),
	}

	err = url.Create()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create URL",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"shortUrl": shortUrl,
		"id":       url.ID,
	})

}

func (URLController) DeleteURL(c *gin.Context) {

}

func (URLController) UpdateURL(c *gin.Context) {

}

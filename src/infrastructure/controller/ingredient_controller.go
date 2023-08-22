package controller

import (
	"net/http"
	"welshacademy/src/application"
	"welshacademy/src/domain"

	"github.com/gin-gonic/gin"
)

func IngredientsList(repository domain.IngredientRepository) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(http.StatusOK, repository.FindAll())
	})
}

type PayloadIngredient struct {
	Name string `json:"name`
	Unit string `json:"unit`
}

func CreateIngredient(service application.CreateIngredient) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var requestBody PayloadIngredient
		c.BindJSON(&requestBody)
		ingredient, err := service.Create(
			requestBody.Name,
			requestBody.Unit,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusCreated, ingredient)
	})
}

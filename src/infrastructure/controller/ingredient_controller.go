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

func CreateIngredient(service application.CreateIngredient) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		ingredient, err := service.Create(
			c.PostForm("name"),
			c.PostForm("unit"),
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusCreated, ingredient)
	})
}

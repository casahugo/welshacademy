package controller

import (
	"net/http"
	"strings"
	"welshacademy/src/application"
	"welshacademy/src/domain"

	"github.com/gin-gonic/gin"
)

func RecipesList(repository domain.RecipeRepository) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		ingredients := c.Query("ingredients")
		var ids []string

		if ingredients != "" {
			for _, id := range strings.Split(ingredients, ",") {
				ids = append(ids, id)
			}
		}

		c.JSON(http.StatusOK, repository.Find(ids))
	})
}

type PayloadIngredients struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}

type PayloadRecipe struct {
	Name         string               `json:"name"`
	Duration     int                  `json:"duration"`
	Descriptions []string             `json:"descriptions"`
	Ingredients  []PayloadIngredients `json:"ingredients"`
}

func CreateRecipe(service application.CreateRecipe) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var requestBody PayloadRecipe
		ingredients := make(map[int]int)
		c.BindJSON(&requestBody)

		for _, item := range requestBody.Ingredients {
			ingredients[item.Id] = item.Quantity
		}

		recipe, err := service.Create(
			requestBody.Name,
			requestBody.Descriptions,
			requestBody.Duration,
			ingredients,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusCreated, recipe)
	})
}

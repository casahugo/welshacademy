package controller

import (
	"net/http"
	"strconv"
	"strings"
	"welshacademy/src/application"
	"welshacademy/src/domain"

	"github.com/gin-gonic/gin"
)

func RecipesList(repository domain.RecipeRepository) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		ingredients := c.PostForm("ingredients")

		var ids []int

		if ingredients != "" {
			for _, split := range strings.Split("ingredients", ",") {
				id, _ := strconv.Atoi(split)
				ids = append(ids, id)
			}
		}

		c.JSON(http.StatusOK, repository.Find(ids))
	})
}

type PayloadRecipe struct {
	Name         string      `json:"name"`
	Duration     int         `json:"duration"`
	Descriptions []string    `json:"descriptions"`
	Ingredients  map[int]int `json:"ingredients"`
}

func CreateRecipe(service application.CreateRecipe) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var requestBody PayloadRecipe

		c.BindJSON(&requestBody)

		recipe, err := service.Create(
			requestBody.Name,
			requestBody.Descriptions,
			requestBody.Duration,
			requestBody.Ingredients,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusCreated, recipe)
	})
}

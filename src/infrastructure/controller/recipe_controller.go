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

func CreateRecipe(service application.CreateRecipe) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		recipe, err := service.Create(
			c.PostForm("name"),
			c.PostFormArray("description"),
			c.GetInt("duration"),
			map[int]int{},
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, recipe)
	})
}

package controller

import (
	"net/http"
	"strconv"
	"welshacademy/src/application"
	"welshacademy/src/domain"

	"github.com/gin-gonic/gin"
)

func FavoriteList(repository domain.FavoriteRepository) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Retrieve userId from session or jwt token
		user := domain.User{Id: 1}

		c.JSON(http.StatusOK, repository.FindByUserId(user.Id))
	})
}

func AddFavorite(service application.FlagFavoriteRecipe) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Retrieve user from session or jwt token
		user := domain.User{Id: 1}
		recipeId, _ := strconv.Atoi(c.Param("recipeId"))

		err := service.Add(user, recipeId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusCreated, nil)
	})
}

func RemoveFavorite(service application.FlagFavoriteRecipe) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Retrieve user from session or jwt token
		user := domain.User{Id: 1}
		recipeId, _ := strconv.Atoi(c.Param("recipeId"))

		err := service.Remove(user, recipeId)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusNoContent, nil)
	})
}

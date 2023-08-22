package main

import (
	"net/http"

	kernel "welshacademy/src"
	"welshacademy/src/infrastructure/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter(application kernel.App) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welsh Academy api")
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ingredient", controller.IngredientsList(application.IngredientRepository))
		v1.POST("/ingredient", controller.CreateIngredient(application.CreateIngredient))

		v1.GET("/recipe", controller.RecipesList(application.RecipeRepository))
		v1.POST("/recipe", controller.CreateRecipe(application.CreateRecipe))

		v1.GET("/favorite", controller.FavoriteList(application.FavoriteRepository))
		v1.POST("/favorite", controller.AddFavorite(application.FlagFavoriteRecipe))
		v1.DELETE("/favorite", controller.RemoveFavorite(application.FlagFavoriteRecipe))
	}

	return router
}

func main() {
	application, _ := kernel.Boot()

	router := setupRouter(*application)

	http.ListenAndServe(":80", router)
}

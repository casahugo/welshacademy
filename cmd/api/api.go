package main

import (
	"net/http"

	kernel "welshacademy/src"
	"welshacademy/src/infrastructure/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter(app kernel.App) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welsh Academy api")
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ingredient", controller.IngredientsList(app.IngredientRepository))
		v1.POST("/ingredient", controller.CreateIngredient(app.CreateIngredient))

		v1.GET("/recipe", controller.RecipesList(app.RecipeRepository))
		v1.POST("/recipe", controller.CreateRecipe(app.CreateRecipe))

		v1.GET("/favorite", controller.FavoriteList(app.FavoriteRepository))
		v1.POST("/favorite/:recipeId", controller.AddFavorite(app.FlagFavoriteRecipe))
		v1.DELETE("/favorite/:recipeId", controller.RemoveFavorite(app.FlagFavoriteRecipe))
	}

	return router
}

func main() {
	application, _ := kernel.Boot()

	router := setupRouter(*application)

	http.ListenAndServe(":80", router)
}

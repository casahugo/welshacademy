package domain

import (
	"time"
)

type QuantityIngredient struct {
	Ingredient Ingredient `json:"ingredient"`
	Quantity   int        `json:"quantity"`
}

type Recipe struct {
	Id          int                  `json:"id"`
	Name        string               `json:"name"`
	Ingredients []QuantityIngredient `json:"ingredients"`
	Description []string             `json:"description"`
	Duration    time.Duration        `json:"duration"`
}

func (r *Recipe) AddIngredient(ingredient Ingredient, quantity int) {
	r.Ingredients = append(r.Ingredients, QuantityIngredient{Ingredient: ingredient, Quantity: quantity})
}

type Recipes []Recipe

type RecipeNotFound struct{}

func (e *RecipeNotFound) Error() string {
	return "recipe not found"
}

type RecipeRepository interface {
	Find(filters []string) Recipes
	Get(id int) (Recipe, *RecipeNotFound)
	Save(entity Recipe) (Recipe, error)
}

type InMemoryRecipeRepository struct {
	Recipes Recipes
}

func (s InMemoryRecipeRepository) Find(filters []string) Recipes {
	return Recipes{}
}

func (s InMemoryRecipeRepository) Get(id int) (Recipe, *RecipeNotFound) {
	for _, recipe := range s.Recipes {
		if recipe.Id == id {
			return recipe, nil
		}
	}

	return Recipe{}, &RecipeNotFound{}
}

func (s InMemoryRecipeRepository) Save(entity Recipe) (Recipe, error) {
	return entity, nil
}

package domain

import (
	"time"
)

type QuantityIngredient struct {
	Ingredient Ingredient
	Quantity   int
}

type Recipe struct {
	Id          int
	Name        string
	Ingredients []QuantityIngredient
	Description []string
	Duration    time.Duration
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
	FindAll(filters string) (Recipes, error)
	Get(id int) (Recipe, *RecipeNotFound)
	Save(entity Recipe) error
}

type InMemoryRecipeRepository struct {
	Recipes Recipes
}

func (s InMemoryRecipeRepository) FindAll(filters string) (Recipes, error) {
	return Recipes{}, nil
}

func (s InMemoryRecipeRepository) Get(id int) (Recipe, *RecipeNotFound) {
	for _, recipe := range s.Recipes {
		if recipe.Id == id {
			return recipe, nil
		}
	}

	return Recipe{}, &RecipeNotFound{}
}

func (s InMemoryRecipeRepository) Save(entity Recipe) error {
	return nil
}

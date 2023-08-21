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

func (r *Recipe) AddIngredient(ingredient Ingredient, quantity int) error {
	r.Ingredients = append(r.Ingredients, QuantityIngredient{Ingredient: ingredient, Quantity: quantity})

	return nil
}

type Recipes []Recipe

type RecipeRepository interface {
	FindAll(filters string) (Recipes, error)
	Save(entity Recipe) error
}

type InMemoryRecipeRepository struct {
}

func (s InMemoryRecipeRepository) FindAll(filters string) (Recipes, error) {
	return Recipes{}, nil
}

func (s InMemoryRecipeRepository) Save(entity Recipe) error {
	return nil
}

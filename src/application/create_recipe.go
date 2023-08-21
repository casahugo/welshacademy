package application

import (
	"errors"
	"time"
	"welshacademy/src/domain"
)

type CreateRecipe struct {
	Repository           domain.RecipeRepository
	IngredientRepository domain.IngredientRepository
}

func (service CreateRecipe) Create(
	name string,
	description []string,
	duration int,
	ingredients map[int]int,
) (domain.Recipe, error) {
	recipe := domain.Recipe{
		Name:        name,
		Duration:    time.Duration(duration) * time.Minute,
		Description: description,
	}

	for id, quantity := range ingredients {
		ingredient, err := service.IngredientRepository.Get(id)

		if err != nil {
			return recipe, errors.New("ingredient not found")
		}

		err = recipe.AddIngredient(ingredient, quantity)

		if err != nil {
			return recipe, err
		}
	}

	return recipe, nil
}

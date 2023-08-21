package application

import "welshacademy/src/domain"

type FlagFavoriteRecipe struct {
	UserRepository   domain.UserRepository
	RecipeRepository domain.RecipeRepository
}

func (service FlagFavoriteRecipe) Add(recipeId int, userId int) error {
	user, err := service.UserRepository.Get(userId)

	if err != nil {
		return err
	}

	recipe, err := service.RecipeRepository.Get(recipeId)

	if err != nil {
		return err
	}

	user.AddFavorite(recipe)

	service.UserRepository.Save(user)

	return nil
}

func (service FlagFavoriteRecipe) Remove(recipeId int, userId int) error {
	user, err := service.UserRepository.Get(userId)

	if err != nil {
		return err
	}

	recipe, err := service.RecipeRepository.Get(recipeId)

	if err != nil {
		return err
	}

	user.RemoveFavorite(recipe)

	service.UserRepository.Save(user)

	return nil
}

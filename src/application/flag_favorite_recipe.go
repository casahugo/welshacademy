package application

import "welshacademy/src/domain"

type FlagFavoriteRecipe struct {
	FavoriteRepository domain.FavoriteRepository
	RecipeRepository   domain.RecipeRepository
}

func (service FlagFavoriteRecipe) Add(user domain.User, recipeId int) error {
	recipe, err := service.RecipeRepository.Get(recipeId)

	if err != nil {
		return err
	}

	favorite := domain.Favorite{UserId: user.Id, RecipeId: recipe.Id}

	service.FavoriteRepository.Save(favorite)

	return nil
}

func (service FlagFavoriteRecipe) Remove(user domain.User, recipeId int) error {
	favorite, err := service.FavoriteRepository.Get(user.Id, recipeId)

	if err != nil {
		return err
	}

	service.FavoriteRepository.Delete(favorite)

	return nil
}

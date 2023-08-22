package application

import (
	"testing"
	"welshacademy/src/domain"

	"github.com/magiconair/properties/assert"
)

func TestAddFavorite(t *testing.T) {
	service := FlagFavoriteRecipe{
		FavoriteRepository: domain.InMemoryFavoriteRepository{},
		RecipeRepository: domain.InMemoryRecipeRepository{
			Recipes: domain.Recipes{
				domain.Recipe{Id: 1, Name: "Welsh aux pommes", Duration: 20},
			},
		},
	}
	user := domain.User{Id: 5, Login: "john.doe"}

	err := service.Add(user, 1)

	assert.Equal(t, err, nil)
}

func TestAddFavoriteNotFound(t *testing.T) {
	service := FlagFavoriteRecipe{
		FavoriteRepository: domain.InMemoryFavoriteRepository{},
		RecipeRepository: domain.InMemoryRecipeRepository{
			Recipes: domain.Recipes{
				domain.Recipe{Id: 1, Name: "Welsh aux pommes", Duration: 20},
			},
		},
	}
	user := domain.User{Id: 5, Login: "john.doe"}

	err := service.Add(user, 10)

	assert.Equal(t, err.Error(), "recipe not found")
}

func TestRemoveFavorite(t *testing.T) {
	service := FlagFavoriteRecipe{
		FavoriteRepository: domain.InMemoryFavoriteRepository{
			Favorites: []domain.Favorite{
				domain.Favorite{UserId: 5, RecipeId: 1},
			},
		},
		RecipeRepository: domain.InMemoryRecipeRepository{},
	}

	err := service.Remove(domain.User{Id: 5, Login: "john.doe"}, 1)

	assert.Equal(t, err, nil)
}

func TestRemoveFavoriteNotFound(t *testing.T) {
	service := FlagFavoriteRecipe{
		FavoriteRepository: domain.InMemoryFavoriteRepository{
			Favorites: []domain.Favorite{
				domain.Favorite{UserId: 1, RecipeId: 1},
			},
		},
		RecipeRepository: domain.InMemoryRecipeRepository{},
	}

	err := service.Remove(domain.User{Id: 5, Login: "john.doe"}, 1)

	assert.Equal(t, err.Error(), "favorite not found")
}

package domain

import "errors"

type User struct {
	Id    int
	Login string
}

type Favorite struct {
	UserId   int
	RecipeId int
}

type FavoriteRepository interface {
	Get(userId int, recipeId int) (Favorite, error)
	FindByUserId(id int) Recipes
	Save(favorite Favorite) error
	Delete(Favorite Favorite) error
}

type InMemoryFavoriteRepository struct {
	Favorites []Favorite
}

func (s InMemoryFavoriteRepository) Get(userId int, recipeId int) (Favorite, error) {
	for _, favorite := range s.Favorites {
		if favorite.UserId == userId && favorite.RecipeId == recipeId {
			return favorite, nil
		}
	}

	return Favorite{}, errors.New("favorite not found")
}

func (s InMemoryFavoriteRepository) FindByUserId(id int) Recipes {
	return Recipes{}
}

func (s InMemoryFavoriteRepository) Save(favorite Favorite) error {
	return nil
}

func (s InMemoryFavoriteRepository) Delete(Favorite Favorite) error {
	return nil
}

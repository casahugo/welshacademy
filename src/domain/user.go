package domain

type User struct {
	Id              int
	Login           string
	FavoriteRecipes []Recipe
}

func (u *User) AddFavorite(recipe Recipe) {
	u.FavoriteRecipes = append(u.FavoriteRecipes, recipe)
}

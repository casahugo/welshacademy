package domain

type User struct {
	Id              int
	Login           string
	FavoriteRecipes []Recipe
}

func (u *User) AddFavorite(recipe Recipe) {
	u.FavoriteRecipes = append(u.FavoriteRecipes, recipe)
}

func (u *User) RemoveFavorite(recipe Recipe) {
}

type UserRepository interface {
	Get(id int) (User, error)
	Save(user User) error
}

type InMemoryUserRepository struct {
	User User
}

func (r InMemoryUserRepository) Get(id int) (User, error) {
	return r.User, nil
}

func (r InMemoryUserRepository) Save(user User) error {
	return nil
}

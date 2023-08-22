package mysql

import (
	"database/sql"
	"welshacademy/src/domain"
)

type FavoriteRepository struct {
	DBMysql *sql.DB
}

func (r FavoriteRepository) Get(userId int, recipeId int) (domain.Favorite, *domain.FavoriteNotFound) {

	return domain.Favorite{}, &domain.FavoriteNotFound{}
}

func (r FavoriteRepository) FindByUserId(id int) domain.Recipes {
	var recipes domain.Recipes

	rows, _ := r.DBMysql.Query("select recipe.* from favorite inner join recipe on recipe.id = favorite.recipe_id where favorite.user_id = ?", id)

	for rows.Next() {
		var recipe domain.Recipe
		rows.Scan(
			&recipe.Id,
			&recipe.Name,
			&recipe.Duration,
		)

		recipes = append(recipes, recipe)
	}

	return recipes
}

func (r FavoriteRepository) Save(favorite domain.Favorite) error {
	return nil
}

func (r FavoriteRepository) Delete(Favorite domain.Favorite) error {
	return nil
}

package mysql

import (
	"database/sql"
	"welshacademy/src/domain"
)

type FavoriteRepository struct {
	DBMysql *sql.DB
}

func (r FavoriteRepository) Get(userId int, recipeId int) (domain.Favorite, *domain.FavoriteNotFound) {
	var favorite domain.Favorite

	err := r.DBMysql.QueryRow("select * from favorite where user_id = ? and recipe_id = ?", userId, recipeId).Scan(
		&favorite.UserId,
		&favorite.RecipeId,
	)

	if err != nil {
		return domain.Favorite{}, &domain.FavoriteNotFound{}
	}

	return favorite, nil
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

		rowsIngredients, _ := r.DBMysql.Query("select ingredient.id, ingredient.name, ingredient.unit, recipe_ingredient.quantity FROM recipe_ingredient inner join ingredient on ingredient.id = recipe_ingredient.ingredient_id where recipe_ingredient.recipe_id = ?", recipe.Id)

		for rowsIngredients.Next() {
			var ingredient domain.Ingredient
			var quantity int
			rowsIngredients.Scan(
				&ingredient.Id,
				&ingredient.Name,
				&ingredient.Unit,
				&quantity,
			)

			recipe.Ingredients = append(recipe.Ingredients, domain.QuantityIngredient{Ingredient: ingredient, Quantity: quantity})
		}
		rowsIngredients.Close()

		rowsDesc, _ := r.DBMysql.Query("select description from recipe_description where recipe_id = ? order by id", recipe.Id)

		for rowsDesc.Next() {
			var desc string
			rowsDesc.Scan(&desc)

			recipe.Description = append(recipe.Description, desc)
		}
		rowsDesc.Close()

		recipes = append(recipes, recipe)
	}

	return recipes
}

func (r FavoriteRepository) Save(favorite domain.Favorite) error {
	_, err := r.DBMysql.Exec("insert into favorite (user_id, recipe_id) values (?, ?)", favorite.UserId, favorite.RecipeId)

	return err
}

func (r FavoriteRepository) Delete(favorite domain.Favorite) error {
	_, err := r.DBMysql.Exec("delete from favorite where recipe_id = ?", favorite.RecipeId)

	return err
}

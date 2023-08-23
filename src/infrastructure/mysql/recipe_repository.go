package mysql

import (
	"database/sql"
	"fmt"
	"strings"
	"welshacademy/src/domain"
)

type RecipeRepository struct {
	DBMysql *sql.DB
}

func (r RecipeRepository) Find(filters []string) domain.Recipes {
	var recipes domain.Recipes
	var rows *sql.Rows
	fmt.Println(filters)
	if len(filters) > 0 {
		rows, _ = r.DBMysql.Query("select recipe.* FROM recipe left join recipe_ingredient i on recipe.id = i.recipe_id where i.ingredient_id IN (?) group by recipe.id", strings.Join(filters, ","))
	} else {
		rows, _ = r.DBMysql.Query("select * FROM recipe")
	}

	defer rows.Close()

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

func (r RecipeRepository) Get(id int) (domain.Recipe, *domain.RecipeNotFound) {
	var recipe domain.Recipe

	err := r.DBMysql.QueryRow("select * from recipe where id = ?", id).Scan(
		&recipe.Id,
		&recipe.Name,
		&recipe.Duration,
	)

	rowsIngredients, _ := r.DBMysql.Query("select ingredient.*, recipe_ingredient.quantity FROM recipe_ingredient inner join ingredient on ingredient.id = recipe_ingredient.ingredient_id where recipe_ingredient.recipe_id = ?", id)

	for rowsIngredients.Next() {
		var ingredient domain.Ingredient
		rowsIngredients.Scan(
			&ingredient.Id,
			&ingredient.Name,
			&ingredient.Unit,
		)

		recipe.Ingredients = append(recipe.Ingredients, domain.QuantityIngredient{Ingredient: ingredient, Quantity: 1})
	}

	rowsIngredients.Close()

	rowsDesc, _ := r.DBMysql.Query("select * from recipe_description where recipe_id = ? order by id", id)

	for rowsDesc.Next() {
		var desc string
		rowsDesc.Scan(&desc)

		recipe.Description = append(recipe.Description, desc)
	}

	rowsDesc.Close()

	if err != nil {
		return domain.Recipe{}, &domain.RecipeNotFound{}
	}

	return recipe, nil
}

func (r RecipeRepository) Save(entity domain.Recipe) error {
	if entity.Id == 0 {
		result, err := r.DBMysql.Exec(
			"insert into recipe (name, duration) VALUES (?, ?)",
			entity.Name,
			int(entity.Duration.Minutes()),
		)

		if err != nil {
			return err
		}

		id, _ := result.LastInsertId()

		entity.Id = int(id)

		for _, desc := range entity.Description {
			r.DBMysql.Exec(
				"insert into recipe_description (recipe_id, description) VALUES (?, ?)",
				entity.Id,
				desc,
			)
		}

		for _, ingredient := range entity.Ingredients {
			r.DBMysql.Exec(
				"insert into recipe_ingredient (recipe_id, ingredient_id, quantity) VALUES (?, ?, ?)",
				entity.Id,
				ingredient.Ingredient.Id,
				ingredient.Quantity,
			)
		}
	}

	return nil
}

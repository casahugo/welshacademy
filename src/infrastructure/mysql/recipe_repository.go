package mysql

import (
	"database/sql"
	"welshacademy/src/domain"
)

type RecipeRepository struct {
	DBMysql *sql.DB
}

func (r RecipeRepository) Find(filters []int) domain.Recipes {
	var recipes domain.Recipes

	rows, _ := r.DBMysql.Query("select * FROM recipe")

	defer rows.Close()

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

func (r RecipeRepository) Get(id int) (domain.Recipe, *domain.RecipeNotFound) {
	var recipe domain.Recipe

	err := r.DBMysql.QueryRow("select * from recipe where id = ?", id).Scan(
		&recipe.Id,
		&recipe.Name,
		&recipe.Duration,
	)

	if err != nil {
		return domain.Recipe{}, &domain.RecipeNotFound{}
	}

	return recipe, nil
}

func (r RecipeRepository) Save(entity domain.Recipe) error {
	if entity.Id > 0 {
		result, err := r.DBMysql.Exec(
			"insert into recipe (name, duration) VALUES (?, ?)",
			entity.Name,
			entity.Duration,
		)

		if err != nil {
			return err
		}

		id, _ := result.LastInsertId()

		entity.Id = int(id)
	}

	return nil
}

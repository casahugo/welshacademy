package mysql

import (
	"database/sql"
	"welshacademy/src/domain"
)

type IngredientRepository struct {
	DBMysql *sql.DB
}

func (r IngredientRepository) Get(id int) (domain.Ingredient, *domain.IngredientNotFound) {
	var ingredient domain.Ingredient

	err := r.DBMysql.QueryRow("select * from ingredient where id = ?", id).Scan(
		&ingredient.Id,
		&ingredient.Name,
		&ingredient.Unit,
	)

	if err != nil {
		return domain.Ingredient{}, &domain.IngredientNotFound{}
	}

	return ingredient, nil
}

func (r IngredientRepository) FindAll() []domain.Ingredient {
	var ingredients []domain.Ingredient
	rows, _ := r.DBMysql.Query("select * FROM ingredient")

	defer rows.Close()

	for rows.Next() {
		var ingredient domain.Ingredient
		rows.Scan(
			&ingredient.Id,
			&ingredient.Name,
			&ingredient.Unit,
		)

		ingredients = append(ingredients, ingredient)
	}

	return ingredients
}

func (r IngredientRepository) Save(entity domain.Ingredient) error {
	if entity.Id == 0 {
		result, err := r.DBMysql.Exec(
			"insert into ingredient (name, unit) VALUES (?, ?)",
			entity.Name,
			entity.Unit,
		)

		if err != nil {
			return err
		}

		id, _ := result.LastInsertId()

		entity.Id = int(id)
	}

	return nil
}

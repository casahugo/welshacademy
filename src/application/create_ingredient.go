package application

import "welshacademy/src/domain"

type CreateIngredient struct {
	Repository domain.IngredientRepository
}

func (service CreateIngredient) Create(nom string, unit string) (*domain.Ingredient, error) {
	u := domain.Unit(unit)

	if err := u.IsValid(); err != nil {
		return nil, err
	}

	ingredient := domain.Ingredient{
		Name: nom,
		Unit: u,
	}

	error := service.Repository.Save(ingredient)

	return &ingredient, error
}

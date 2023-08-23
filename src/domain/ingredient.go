package domain

import "errors"

type Unit string

const (
	Undefined      Unit = ""
	Gramme         Unit = "g"
	Centilitre     Unit = "cl"
	CuilliereSoupe Unit = "c.à.s"
	Tranche        Unit = "tranche"
)

func (u Unit) IsValid() error {
	switch u {
	case Undefined, Gramme, Centilitre, CuilliereSoupe, Tranche:
		return nil
	}
	return errors.New("invalid unit type")
}

type Ingredient struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Unit Unit   `json:"unit"`
}

type IngredientNotFound struct{}

func (e *IngredientNotFound) Error() string {
	return "ingredient not found"
}

type IngredientRepository interface {
	Get(id int) (Ingredient, *IngredientNotFound)
	FindAll() []Ingredient
	Save(entity Ingredient) (Ingredient, error)
}

type InMemoryIngredientRepository struct {
	Data []Ingredient
}

func (r InMemoryIngredientRepository) Get(id int) (Ingredient, *IngredientNotFound) {
	for _, ingredient := range r.Data {
		if ingredient.Id == id {
			return ingredient, nil
		}
	}

	return Ingredient{}, &IngredientNotFound{}
}

func (r InMemoryIngredientRepository) FindAll() []Ingredient {
	return []Ingredient{}
}

func (r InMemoryIngredientRepository) Save(entity Ingredient) (Ingredient, error) {
	return entity, nil
}

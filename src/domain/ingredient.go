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
	Id   int
	Name string
	Unit Unit
}

type IngredientRepository interface {
	Get(id int) (Ingredient, error)
	FindAll() ([]Ingredient, error)
	Save(entity Ingredient) error
}

type InMemoryIngredientRepository struct {
	Data []Ingredient
}

func (r InMemoryIngredientRepository) Get(id int) (Ingredient, error) {
	for _, ingredient := range r.Data {
		if ingredient.Id == id {
			return ingredient, nil
		}
	}

	return Ingredient{}, errors.New("ingredient not found")
}

func (r InMemoryIngredientRepository) FindAll() ([]Ingredient, error) {
	return []Ingredient{}, nil
}

func (r InMemoryIngredientRepository) Save(entity Ingredient) error {
	return nil
}

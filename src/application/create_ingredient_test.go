package application

import (
	"testing"
	"welshacademy/src/domain"

	"github.com/magiconair/properties/assert"
)

func TestCreateIngredient(t *testing.T) {
	service := CreateIngredient{
		Repository: domain.InMemoryIngredientRepository{},
	}

	ingredient, _ := service.Create("cheddar", "g")

	assert.Equal(t, ingredient.Name, "cheddar")
	assert.Equal(t, ingredient.Unit, domain.Gramme)
}

func TestCreateIngredientWithUndefinedUnit(t *testing.T) {
	service := CreateIngredient{
		Repository: domain.InMemoryIngredientRepository{},
	}

	ingredient, _ := service.Create("poivre", "")

	assert.Equal(t, ingredient.Name, "poivre")
	assert.Equal(t, ingredient.Unit, domain.Undefined)
}

func TestCreateIngredientWithInvalidUnit(t *testing.T) {
	service := CreateIngredient{
		Repository: domain.InMemoryIngredientRepository{},
	}

	_, err := service.Create("cheddar", "gg")

	assert.Equal(t, "invalid unit type", err.Error())
}

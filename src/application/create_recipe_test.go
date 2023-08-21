package application

import (
	"testing"
	"welshacademy/src/domain"

	"github.com/magiconair/properties/assert"
)

func TestCreateRecipe(t *testing.T) {
	service := CreateRecipe{
		Repository: domain.InMemoryRecipeRepository{},
		IngredientRepository: domain.InMemoryIngredientRepository{
			[]domain.Ingredient{
				domain.Ingredient{Id: 1, Name: "Bière brune", Unit: "cl"},
				domain.Ingredient{Id: 2, Name: "Cheddar", Unit: "g"},
			},
		},
	}

	recipe, _ := service.Create(
		"Welsh traditionnel à la bière brune",
		[]string{
			"Couper le cheddar en petits cubes. Couper le pain en tranches bien épaisses. Faire en sorte qu’au niveau largeur elles passent dans les plats à welsh. Dans la limite du possible garder la croute des tartines.",
			"Verser quelques goutes de bière sur chaque tartine (vraiment quelques gouttes, il faut garder environ 20 cl pour le reste de la recette). Puis les badigeonner d’un peu de moutarde (environ 2 cuillères à soupe) et les placer dans le fond des plats.",
		},
		40,
		map[int]int{1: 4, 2: 800},
	)

	assert.Equal(t, recipe.Name, "Welsh traditionnel à la bière brune")
	assert.Equal(t, recipe.Description[0], "Couper le cheddar en petits cubes. Couper le pain en tranches bien épaisses. Faire en sorte qu’au niveau largeur elles passent dans les plats à welsh. Dans la limite du possible garder la croute des tartines.")
	assert.Equal(t, recipe.Ingredients[0], domain.QuantityIngredient{
		Ingredient: domain.Ingredient{Id: 1, Name: "Bière brune", Unit: "cl"},
		Quantity:   4,
	})
	assert.Equal(t, recipe.Ingredients[1], domain.QuantityIngredient{
		Ingredient: domain.Ingredient{Id: 2, Name: "Cheddar", Unit: "g"},
		Quantity:   800,
	})
}

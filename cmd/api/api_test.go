package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	kernel "welshacademy/src"
	"welshacademy/src/infrastructure/controller"

	"github.com/magiconair/properties/assert"
)

func TestIngredientList(t *testing.T) {
	os.Setenv("DATABASE_URL", "root@(db:3306)/welsh")
	application, _ := kernel.Boot()
	router := setupRouter(*application)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ingredient", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(
		t,
		`[{"id":1,"name":"mourtarde à l'ancienne","unit":"c.à.s"},{"id":2,"name":"poivre","unit":""},{"id":3,"name":"cheddar","unit":"g"},{"id":4,"name":"bière brune","unit":"cl"},{"id":5,"name":"pain de campagne non tranchée","unit":""},{"id":6,"name":"jambon blanc","unit":"tranche"},{"id":7,"name":"oeufs","unit":""},{"id":8,"name":"bière blonde","unit":"cl"},{"id":9,"name":"saint-marcelin","unit":"g"}]`,
		w.Body.String(),
	)
}

func TestCreateIngredient(t *testing.T) {
	os.Setenv("DATABASE_URL", "root@(db:3306)/welsh")
	application, _ := kernel.Boot()
	router := setupRouter(*application)

	w := httptest.NewRecorder()

	jsonValue, _ := json.Marshal(controller.PayloadIngredient{
		Name: "vin blanc",
		Unit: "cl",
	})

	req, _ := http.NewRequest("POST", "/api/v1/ingredient", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestRecipeList(t *testing.T) {
	os.Setenv("DATABASE_URL", "root@(db:3306)/welsh")
	application, _ := kernel.Boot()
	router := setupRouter(*application)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/recipe", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `[{"id":1,"name":"Welsh traditionnel à la bière brune","ingredients":[{"ingredient":{"id":1,"name":"mourtarde à l'ancienne","unit":"c.à.s"},"quantity":4},{"ingredient":{"id":2,"name":"poivre","unit":""},"quantity":1},{"ingredient":{"id":3,"name":"cheddar","unit":"g"},"quantity":800},{"ingredient":{"id":4,"name":"bière brune","unit":"cl"},"quantity":25},{"ingredient":{"id":5,"name":"pain de campagne non tranchée","unit":""},"quantity":25},{"ingredient":{"id":6,"name":"jambon blanc","unit":"tranche"},"quantity":1},{"ingredient":{"id":7,"name":"oeufs","unit":""},"quantity":4}],"description":["Couper le cheddar en petits cubes. Couper le pain en tranches bien épaisses. Faire en sorte qu’au niveau largeur elles passent dans les plats à welsh. Dans la limite du possible garder la croute des tartines.","Verser quelques goutes de bière sur chaque tartine (vraiment quelques gouttes, il faut garder environ 20 cl pour le reste de la recette). Puis les badigeonner d’un peu de moutarde (environ 2 cuillères à soupe) et les placer dans le fond des plats.","Poser sur chaque tartine une tranche de jambon recourbée sur elle-même.","Dans une grande sauteuse, faire fondre les cubes de cheddar sans ajouter de matière grasse.","Mélanger constamment à l’aide d’une cuillère en bois (ça fait des fils, c’est normal tant que ça n’attache pas à la sauteuse).","Quand la consistance est relativement homogène et qu’elle recouvre bien la cuillère quand on mélange, ajouter toute la bière et continuer à mélanger. Il faut remuer sans arrêt de façon à bien incorporer la bière au fromage.","Une fois le tout bien homogène et onctueux, ajouter le reste de la moutarde, un peu de poivre et remuer encore.","Quand la préparation est homogène, la verser dans les plats pour napper les tartines au jambon. Enfourner à 210°C (thermostat 7) pour une 10-12 minutes (il faut que ce soit bien doré).","Pendant ce temps cuire les 4 oeufs au plat dans une poêle (il est aussi possible de les faire au four, directement sur les Welshs, mais la cuisson du jaune est très délicate. C'est pourquoi à la poêle c'est très bien).","Quand le fromage est bien doré, sortir les plats du four, placer un oeuf au plat sur le dessus de chacun d’entre eux, tourner un coup de moulin à poivre et servir rapidement avec des frites et/ou de la salade verte. Et surtout une bière !"],"duration":40},{"id":2,"name":"Welsh saint-marcelin à la bière blonde","ingredients":[{"ingredient":{"id":1,"name":"mourtarde à l'ancienne","unit":"c.à.s"},"quantity":4},{"ingredient":{"id":2,"name":"poivre","unit":""},"quantity":1},{"ingredient":{"id":5,"name":"pain de campagne non tranchée","unit":""},"quantity":25},{"ingredient":{"id":6,"name":"jambon blanc","unit":"tranche"},"quantity":1},{"ingredient":{"id":7,"name":"oeufs","unit":""},"quantity":4},{"ingredient":{"id":8,"name":"bière blonde","unit":"cl"},"quantity":25},{"ingredient":{"id":9,"name":"saint-marcelin","unit":"g"},"quantity":800}],"description":null,"duration":40}]`, w.Body.String())
}

func TestFavoriteRecipeList(t *testing.T) {
	os.Setenv("DATABASE_URL", "root@(db:3306)/welsh")
	application, _ := kernel.Boot()
	router := setupRouter(*application)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/favorite", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `[{"id":1,"name":"Welsh traditionnel à la bière brune","ingredients":[{"ingredient":{"id":1,"name":"mourtarde à l'ancienne","unit":"c.à.s"},"quantity":4},{"ingredient":{"id":2,"name":"poivre","unit":""},"quantity":1},{"ingredient":{"id":3,"name":"cheddar","unit":"g"},"quantity":800},{"ingredient":{"id":4,"name":"bière brune","unit":"cl"},"quantity":25},{"ingredient":{"id":5,"name":"pain de campagne non tranchée","unit":""},"quantity":25},{"ingredient":{"id":6,"name":"jambon blanc","unit":"tranche"},"quantity":1},{"ingredient":{"id":7,"name":"oeufs","unit":""},"quantity":4}],"description":["Couper le cheddar en petits cubes. Couper le pain en tranches bien épaisses. Faire en sorte qu’au niveau largeur elles passent dans les plats à welsh. Dans la limite du possible garder la croute des tartines.","Verser quelques goutes de bière sur chaque tartine (vraiment quelques gouttes, il faut garder environ 20 cl pour le reste de la recette). Puis les badigeonner d’un peu de moutarde (environ 2 cuillères à soupe) et les placer dans le fond des plats.","Poser sur chaque tartine une tranche de jambon recourbée sur elle-même.","Dans une grande sauteuse, faire fondre les cubes de cheddar sans ajouter de matière grasse.","Mélanger constamment à l’aide d’une cuillère en bois (ça fait des fils, c’est normal tant que ça n’attache pas à la sauteuse).","Quand la consistance est relativement homogène et qu’elle recouvre bien la cuillère quand on mélange, ajouter toute la bière et continuer à mélanger. Il faut remuer sans arrêt de façon à bien incorporer la bière au fromage.","Une fois le tout bien homogène et onctueux, ajouter le reste de la moutarde, un peu de poivre et remuer encore.","Quand la préparation est homogène, la verser dans les plats pour napper les tartines au jambon. Enfourner à 210°C (thermostat 7) pour une 10-12 minutes (il faut que ce soit bien doré).","Pendant ce temps cuire les 4 oeufs au plat dans une poêle (il est aussi possible de les faire au four, directement sur les Welshs, mais la cuisson du jaune est très délicate. C'est pourquoi à la poêle c'est très bien).","Quand le fromage est bien doré, sortir les plats du four, placer un oeuf au plat sur le dessus de chacun d’entre eux, tourner un coup de moulin à poivre et servir rapidement avec des frites et/ou de la salade verte. Et surtout une bière !"],"duration":40}]`, w.Body.String())
}

func TestCreateRecipe(t *testing.T) {
	os.Setenv("DATABASE_URL", "root@(db:3306)/welsh")
	application, _ := kernel.Boot()
	router := setupRouter(*application)

	w := httptest.NewRecorder()

	jsonValue, _ := json.Marshal(controller.PayloadRecipe{
		Name:         "Weslh pied noir",
		Descriptions: []string{"Couper le fromage en rondelle", "Faire préchauffer le four"},
		Duration:     20,
		Ingredients:  map[int]int{1: 2, 3: 100},
	})

	req, _ := http.NewRequest("POST", "/api/v1/recipe", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestAddRemoveFavoriteRecipe(t *testing.T) {
	os.Setenv("DATABASE_URL", "root@(db:3306)/welsh")
	application, _ := kernel.Boot()
	router := setupRouter(*application)

	w2 := httptest.NewRecorder()

	req2, _ := http.NewRequest("DELETE", "/api/v1/favorite/1", nil)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 204, w2.Code)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/v1/favorite/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

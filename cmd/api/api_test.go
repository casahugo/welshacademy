package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	kernel "welshacademy/src"

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
	assert.Equal(t, `[{"id":1,"name":"mourtarde à l'ancienne","unit":"c.à.s"},{"id":2,"name":"poivre","unit":""},{"id":3,"name":"cheddar","unit":"g"},{"id":4,"name":"bière brune","unit":"cl"},{"id":5,"name":"pain de campagne non tranchée","unit":""},{"id":6,"name":"jambon blanc","unit":"tranche"},{"id":7,"name":"oeufs","unit":""}]`, w.Body.String())
}

func TestCreateIngredient(t *testing.T) {
	os.Setenv("DATABASE_URL", "root@(db:3306)/welsh")
	application, _ := kernel.Boot()
	router := setupRouter(*application)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/ingredient", nil)
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
	assert.Equal(t, `[{"Id":1,"Name":"Welsh traditionnel à la bière brune","Ingredients":null,"Description":null,"Duration":40}]`, w.Body.String())
}

func TestFavoriteRecipeList(t *testing.T) {
	os.Setenv("DATABASE_URL", "root@(db:3306)/welsh")
	application, _ := kernel.Boot()
	router := setupRouter(*application)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/favorite", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `[{"Id":1,"Name":"Welsh traditionnel à la bière brune","Ingredients":null,"Description":null,"Duration":40}]`, w.Body.String())
}

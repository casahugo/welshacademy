//go:build wireinject
// +build wireinject

package welshacademy

import (
	"database/sql"
	"os"
	"welshacademy/src/application"
	"welshacademy/src/domain"
	"welshacademy/src/infrastructure/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/joho/godotenv"
)

type App struct {
	DBMysql           *sql.DB
	RecipeService     domain.RecipeService
	IngredientService domain.IngredientService
}

func Boot() (*App, error) {
	godotenv.Load()

	return InitApp()
}

func getDBMysql() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL")+"?multiStatements=true")

	if err != nil {
		panic(err)
	}

	return db
}

func InitApp() (*App, error) {
	panic(wire.Build(
		getDBMysql,

		wire.Struct(new(mysql.RecipeRepository), "*"),
		wire.Bind(new(domain.RecipeRepository), new(mysql.RecipeRepository)),
		application.NewRecipeService,

		wire.Struct(new(mysql.IngredientRepository), "*"),
		wire.Bind(new(domain.IngredientRepository), new(mysql.IngredientRepository)),
		application.NewIngredientService,

		wire.Struct(new(App), "*"),
	))
}

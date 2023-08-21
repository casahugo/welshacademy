//go:build wireinject
// +build wireinject

package welshacademy

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/joho/godotenv"
)

type App struct {
	DBMysql *sql.DB
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
		wire.Struct(new(App), "*"),
	))
}

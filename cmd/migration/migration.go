package main

import (
	"fmt"
	"os"
	application "welshacademy/src"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Lancement des migrations...")

	app, _ := application.Boot()

	driver, err := mysql.WithInstance(app.DBMysql, &mysql.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///var/www/html/.migrations/",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = m.Up()

	if err != nil {
		fmt.Println("Migration is up to date.", err)
		return
	}

	fmt.Println("Migration is done.")
}

package main

import (
	"19_Unit-Testing/Praktikum/RESTful-API-Testing/database"
	"19_Unit-Testing/Praktikum/RESTful-API-Testing/routes"
)

func init() {
	database.InitialMigration()
}

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}

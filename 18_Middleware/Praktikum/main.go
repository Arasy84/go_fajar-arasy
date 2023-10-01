package main

import (
	"18_Middleware/Praktikum/configs"
	"18_Middleware/Praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	configs.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}

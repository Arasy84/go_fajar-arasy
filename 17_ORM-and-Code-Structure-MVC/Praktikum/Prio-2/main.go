package main

import (
	"17_ORM-and-Code-Structure-MVC/Praktikum/Prio-2/configs"
	"17_ORM-and-Code-Structure-MVC/Praktikum/Prio-2/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDB()
	configs.InitialMigration()

	e := echo.New()

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)

	e.Start(":8000")
}

package main

import (
	"17_ORM-and-Code-Structure-MVC/Praktikum/Prio-1/2/configs"
	"17_ORM-and-Code-Structure-MVC/Praktikum/Prio-1/2/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDB()
	configs.InitialMigration()

	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.Start(":8000")
}

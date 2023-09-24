package main

import (
	"17_ORM-and-Code-Structure-MVC/Praktikum/Explo/configs"
	"17_ORM-and-Code-Structure-MVC/Praktikum/Explo/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDB()
	configs.InitialMigration()

	e := echo.New()

	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)

	e.Start(":8000")
}

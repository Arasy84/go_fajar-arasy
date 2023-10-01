package routes

import (
	"19_Unit-Testing/Praktikum/RESTful-API-Testing/constants"
	c "19_Unit-Testing/Praktikum/RESTful-API-Testing/controllers"
	m "19_Unit-Testing/Praktikum/RESTful-API-Testing/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)

	e.POST("/auth/login", c.LoginController)

	e.POST("/users", c.CreateUserController)

	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)

	jwtAuth := e.Group("/restricted")
	jwtAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	jwtAuth.GET("/users", c.GetUsersController)
	jwtAuth.GET("/users/:id", c.GetUserController)
	jwtAuth.DELETE("/users/:id", c.DeleteUserController)
	jwtAuth.PUT("/users/:id", c.UpdateUserController)

	jwtAuth.POST("/books", c.CreateBookController)
	jwtAuth.DELETE("/books/:id", c.DeleteBookController)
	jwtAuth.PUT("/books/:id", c.UpdateBookController)

	return e
}
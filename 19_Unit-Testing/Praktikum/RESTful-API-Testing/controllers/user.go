package controllers

import (
	"19_Unit-Testing/Praktikum/RESTful-API-Testing/database"
	"19_Unit-Testing/Praktikum/RESTful-API-Testing/models"
	formatter "19_Unit-Testing/Praktikum/RESTful-API-Testing/formatter"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	DB := database.Connect()
	var users []models.User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	DB := database.Connect()
	id := c.Param("id")
	user := models.User{}

	if err := DB.Where("ID = ?", id).First(&user).Error; err != nil {
		return formatter.NotFoundResponse(c, err)
	}

	return formatter.SuccessResponse(c, user)
}

// create new user
func CreateUserController(c echo.Context) error {
	DB := database.Connect()
	user := models.User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	DB := database.Connect()
	id := c.Param("id")

	DB.Delete(&models.User{}, id)

	return c.JSON(http.StatusOK, formatter.SuccessResponse(c, nil))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	DB := database.Connect()
	id := c.Param("id")
	user := models.User{}

	DB.Where("ID = ?", id).First(&user)

	if user.ID == 0 {
		return formatter.NotFoundResponse(c, nil) 
	}

	payload := models.User{}
	c.Bind(&payload)

	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password
	DB.Save(&user)

	return formatter.SuccessResponse(c, user) 
}

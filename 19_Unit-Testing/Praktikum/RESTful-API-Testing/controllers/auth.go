package controllers

import (
	"net/http"
	"strconv"

	"19_Unit-Testing/Praktikum/RESTful-API-Testing/database"
	"19_Unit-Testing/Praktikum/RESTful-API-Testing/middleware"
	"19_Unit-Testing/Praktikum/RESTful-API-Testing/models"
	"19_Unit-Testing/Praktikum/RESTful-API-Testing/formatter"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	DB := database.Connect()

	user := models.User{}
	c.Bind(&user)

	err := DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		// Use formatter.InternalServerErrorResponse to create the response
		return c.JSON(http.StatusInternalServerError, formatter.InternalServerErrorResponse(map[string]interface{}{"error": err.Error()}))
	}

	token, err := middleware.CreateToken(strconv.Itoa(int(user.ID)), user.Name)
	if err != nil {
		// Use formatter.InternalServerErrorResponse to create the response
		return c.JSON(http.StatusInternalServerError, formatter.InternalServerErrorResponse(map[string]interface{}{"error": err.Error()}))
	}

	responseData := models.UserLogin{}
	responseData.ID = user.ID
	responseData.Name = user.Name
	responseData.Email = user.Email
	responseData.Token = token

	// Use formatter.SuccessResponse to create the success response
	return c.JSON(http.StatusOK, formatter.SuccessResponse(c, responseData))
}

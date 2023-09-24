package controllers

import (
	"net/http"
	"strconv"

	"17_ORM-and-Code-Structure-MVC/Praktikum/Prio-1/2/configs"
	"17_ORM-and-Code-Structure-MVC/Praktikum/Prio-1/2/models"

	"github.com/labstack/echo/v4"
)

// Get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := configs.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// Get user by ID
func GetUserController(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user ID")
	}

	var user models.User

	if err := configs.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// Create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := configs.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// Delete user by ID
func DeleteUserController(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User

	if err := configs.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := configs.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// Update user by ID
func UpdateUserController(c echo.Context) error {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User

	if err := configs.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	if err := configs.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

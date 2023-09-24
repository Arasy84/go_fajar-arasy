package controllers

import (
	"net/http"
	"strconv"

	"17_ORM-and-Code-Structure-MVC/Praktikum/Explo/configs"
	"17_ORM-and-Code-Structure-MVC/Praktikum/Explo/models"

	"github.com/labstack/echo/v4"
)

// Get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := configs.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// Get blog by ID
func GetBlogController(c echo.Context) error {
	blogIDStr := c.Param("id")
	blogID, err := strconv.Atoi(blogIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid blog ID")
	}

	var blog models.Blog

	if err := configs.DB.First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"blog":    blog,
	})
}

// Create new blog
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	if err := configs.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// Delete blog by ID
func DeleteBlogController(c echo.Context) error {
	blogIDStr := c.Param("id")
	blogID, err := strconv.Atoi(blogIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog models.Blog

	if err := configs.DB.First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := configs.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}

// Update blog by ID
func UpdateBlogController(c echo.Context) error {
	blogIDStr := c.Param("id")
	blogID, err := strconv.Atoi(blogIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog models.Blog

	if err := configs.DB.First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&blog)

	if err := configs.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})
}

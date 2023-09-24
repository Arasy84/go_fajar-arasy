package controllers

import (
	"net/http"
	"strconv"

	"17_ORM-and-Code-Structure-MVC/Praktikum/Prio-2/configs"
	"17_ORM-and-Code-Structure-MVC/Praktikum/Prio-2/models"

	"github.com/labstack/echo/v4"
)

// Get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := configs.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// Get book by ID
func GetBookController(c echo.Context) error {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid book ID")
	}

	var book models.Book

	if err := configs.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// Create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := configs.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// Delete book by ID
func DeleteBookController(c echo.Context) error {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book models.Book

	if err := configs.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := configs.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// Update book by ID
func UpdateBookController(c echo.Context) error {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book models.Book

	if err := configs.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&book)

	if err := configs.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

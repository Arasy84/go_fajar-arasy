package formatter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SuccessResponse formats a success response.
func SuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

// NotFoundResponse formats a not found response.
func NotFoundResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"status":  "error",
		"message": "Not found",
		"error":   err.Error(),
	})
}

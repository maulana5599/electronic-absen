package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authentication(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
	})
}

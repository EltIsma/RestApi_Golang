package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c echo.Context, satusCode int, message string) {
	logrus.Error(message)
	c.JSON(satusCode, errorResponse{message})
}

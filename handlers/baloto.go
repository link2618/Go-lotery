package handlers

import (
	// "encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/link2618/Go-lotery/server"
)

type TestResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func Test(s server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		response := TestResponse{
			Message: "Welcome to lotery Go",
			Status:  true,
		}
		return c.JSON(http.StatusOK, response)
	}
}

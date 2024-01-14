package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/link2618/Go-lotery/models"
	"github.com/link2618/Go-lotery/repository"
	"github.com/link2618/Go-lotery/server"
	"github.com/link2618/Go-lotery/validations"
)

type DefaultResponse struct {
	Message string `json:"message"`
	State   string `json:"state"`
}

func Test(s server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		response := DefaultResponse{
			Message: "Welcome to lotery Go",
			State:   "Success",
		}
		return c.JSON(http.StatusOK, response)
	}
}

func InsertNewGame(s server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.Baloto

		if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
			response := DefaultResponse{
				State:   "Error",
				Message: "Ocurrió un error inesperado",
			}
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response)
		}

		// Data is being validated
		msg, isValid := validations.IsValidBaloto(body)
		if isValid == false {
			response := DefaultResponse{
				State:   "Error",
				Message: msg,
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		// We save it in the database
		err := repository.InsertBaloto(c.Request().Context(), &body)
		if err != nil {
			response := DefaultResponse{
				State:   "Error",
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := DefaultResponse{
			Message: "Se creó el registro exitosamente.",
			State:   "Success",
		}
		return c.JSON(http.StatusCreated, response)
	}
}

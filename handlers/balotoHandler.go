package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/link2618/Go-lotery/helpers"
	"github.com/link2618/Go-lotery/models"
	"github.com/link2618/Go-lotery/repository"
	"github.com/link2618/Go-lotery/server"
	"github.com/link2618/Go-lotery/validations"
)

type DefaultResponse struct {
	Message string `json:"message"`
	State   string `json:"state"`
}

type GenerateGameResponse struct {
	State   string `json:"state"`
	Number1 uint8  `json:"number1"`
	Number2 uint8  `json:"number2"`
	Number3 uint8  `json:"number3"`
	Number4 uint8  `json:"number4"`
	Number5 uint8  `json:"number5"`
	Serial  uint8  `json:"serial"`
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
			fmt.Println("Error InsertNewGame", err)
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

func GenerateGame(s server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		var numbers []uint8
		var serie []uint8

		for {
			numbers = helpers.GenerateNumerics(5, 44)
			sort.SliceStable(numbers[:], func(i, j int) bool {
				return numbers[i] < numbers[j]
			})

			serie = helpers.GenerateNumerics(1, 16)

			exists, err := repository.NewGameExists(c.Request().Context(), numbers, serie[0])
			if err != nil {
				response := DefaultResponse{
					State:   "Error",
					Message: err.Error(),
				}
				return c.JSON(http.StatusInternalServerError, response)
			}

			if !exists {
				break
			}
		}

		response := GenerateGameResponse{
			State:   "Success",
			Number1: numbers[0],
			Number2: numbers[1],
			Number3: numbers[2],
			Number4: numbers[3],
			Number5: numbers[4],
			Serial:  serie[0],
		}
		return c.JSON(http.StatusOK, response)
	}
}

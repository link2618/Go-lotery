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

type SearchNumberResponse struct {
	State string `json:"state"`
	Data  []*models.Baloto
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
		msg, isValid := validations.IsValidBaloto(body, true, true)
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

func SearchNumber(s server.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.Baloto

		if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
			response := DefaultResponse{
				State:   "Error",
				Message: "Ocurrió un error inesperado",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		// Data is being validated
		validateSerie := false
		serie := []uint8{}
		if body.Serial != 0 {
			validateSerie = true
			serie = append(serie, body.Serial)
		}
		msg, isValid := validations.IsValidBaloto(body, false, validateSerie)
		if isValid == false {
			response := DefaultResponse{
				State:   "Error",
				Message: msg,
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		numbers := []uint8{body.Number1, body.Number2, body.Number3, body.Number4, body.Number5}

		data, err := repository.SearchNumber(c.Request().Context(), numbers, serie...)
		if err != nil {
			response := DefaultResponse{
				State:   "Error",
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, response)
		}

		response := SearchNumberResponse{
			State: "Success",
			Data:  data,
		}
		return c.JSON(http.StatusOK, response)
	}
}

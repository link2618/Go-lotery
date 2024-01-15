package main

import (
	"context"
	"log"
	"os"

	"github.com/link2618/Go-lotery/handlers"
	"github.com/link2618/Go-lotery/server"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		// panic(err)
		// return
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(Setuphandlers)
}

func Setuphandlers(s server.Server, e *echo.Echo) {
	var pre string = "/api/v1/"

	e.GET(pre+"test", handlers.Test(s))
	e.GET(pre+"generate-game", handlers.GenerateGame(s))
	e.POST(pre+"insert-new-game", handlers.InsertNewGame(s))
	e.POST(pre+"search-number", handlers.SearchNumber(s))
}

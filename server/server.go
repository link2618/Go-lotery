package server

import (
	"context"
	"errors"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/link2618/Go-lotery/database"
	"github.com/link2618/Go-lotery/repository"
	// "github.com/link2618/Go-lotery/handlers"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	echo   *echo.Echo
}

func (b *Broker) Config() *Config {
	return b.config
}

// Constructor
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("Secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("Database is required")
	}

	broker := &Broker{
		config: config,
		echo:   echo.New(),
	}

	return broker, nil
}

// Inicializar el servidor
func (b *Broker) Start(binder func(s Server, e *echo.Echo)) {
	b.echo.Use(middleware.Logger())
	b.echo.Use(middleware.Recover())
	b.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "https://www.web.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	binder(b, b.echo)

	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)

	log.Println("Starting server port", b.config.Port)

	if err := b.echo.Start(b.config.Port); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

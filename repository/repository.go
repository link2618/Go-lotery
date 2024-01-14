package repository

import (
	"context"

	"github.com/link2618/Go-lotery/models"
)

type Repository interface {
	InsertBaloto(ctx context.Context, baloto *models.Baloto) error
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertBaloto(ctx context.Context, baloto *models.Baloto) error {
	return implementation.InsertBaloto(ctx, baloto)
}

func Close() error {
	return implementation.Close()
}

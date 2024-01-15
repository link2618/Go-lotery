package repository

import (
	"context"

	"github.com/link2618/Go-lotery/models"
)

type Repository interface {
	InsertBaloto(ctx context.Context, baloto *models.Baloto) error
	NewGameExists(ctx context.Context, numbers []uint8, serie uint8) (bool, error)
	SearchNumber(ctx context.Context, numbers []uint8, serie ...uint8) ([]*models.Baloto, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertBaloto(ctx context.Context, baloto *models.Baloto) error {
	return implementation.InsertBaloto(ctx, baloto)
}

func NewGameExists(ctx context.Context, numbers []uint8, serie uint8) (bool, error) {
	return implementation.NewGameExists(ctx, numbers, serie)
}

func SearchNumber(ctx context.Context, numbers []uint8, serie ...uint8) ([]*models.Baloto, error) {
	return implementation.SearchNumber(ctx, numbers, serie...)
}

func Close() error {
	return implementation.Close()
}

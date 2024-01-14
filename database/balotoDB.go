package database

import (
	"context"

	"github.com/link2618/Go-lotery/models"
)

func (repo *PostgresRepository) InsertBaloto(ctx context.Context, baloto *models.Baloto) error {
	typeB, number1, number2, number3, number4, number5, serial, date := baloto.Type, baloto.Number1, baloto.Number2, baloto.Number3, baloto.Number4, baloto.Number5, baloto.Serial, baloto.Date
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO lottery.baloto (type, number1, number2, number3, number4, number5, serie, date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		typeB, number1, number2, number3, number4, number5, serial, date,
	)
	return err
}

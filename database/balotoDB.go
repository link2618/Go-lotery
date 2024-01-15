package database

import (
	"context"
	"log"

	_ "github.com/lib/pq"
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

func (repo *PostgresRepository) NewGameExists(ctx context.Context, numbers []uint8, serie uint8) (bool, error) {
	number1, number2, number3, number4, number5 := numbers[0], numbers[1], numbers[2], numbers[3], numbers[4]
	rows, err := repo.db.QueryContext(
		ctx,
		`SELECT *
		FROM lottery.baloto
		WHERE number1 = $1 AND number2 = $2 AND number3 = $3 AND number4 = $4 AND number5 = $5 AND serie = $6;`,
		number1, number2, number3, number4, number5, serie,
	)
	if err != nil {
		return false, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func (repo *PostgresRepository) SearchNumber(ctx context.Context, numbers []uint8, serie ...uint8) ([]*models.Baloto, error) {
	number1, number2, number3, number4, number5 := numbers[0], numbers[1], numbers[2], numbers[3], numbers[4]

	var query string
	var args []interface{}

	if len(serie) > 0 {
		query = `SELECT *
		FROM lottery.baloto
		WHERE number1 = $1 AND number2 = $2 AND number3 = $3 AND number4 = $4 AND number5 = $5 AND serie = $6;`
		args = append(args, number1, number2, number3, number4, number5, serie[0])
	} else {
		query = `SELECT *
		FROM lottery.baloto
		WHERE number1 = $1 AND number2 = $2 AND number3 = $3 AND number4 = $4 AND number5 = $5;`
		args = append(args, number1, number2, number3, number4, number5)
	}

	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var balotos []*models.Baloto
	for rows.Next() {
		var baloto models.Baloto
		if err = rows.Scan(&baloto.Id, &baloto.Type, &baloto.Number1, &baloto.Number2, &baloto.Number3, &baloto.Number4, &baloto.Number5, &baloto.Serial, &baloto.Date); err == nil {
			balotos = append(balotos, &baloto)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return balotos, nil
}

package models

import "time"

// serie del 1 al 16 y numeros del 1 al 44 y no se repiten
type Baloto struct {
	Id      string    `json:"id"`
	Type    string    `json:"type"`
	Number1 uint8     `json:"number1"`
	Number2 uint8     `json:"number2"`
	Number3 uint8     `json:"number3"`
	Number4 uint8     `json:"number4"`
	Number5 uint8     `json:"number5"`
	Serial  uint8     `json:"serial"`
	Date    time.Time `json:"date"`
}

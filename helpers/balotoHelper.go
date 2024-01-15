package helpers

import "math/rand"

func GenerateNumerics(quantityNumbers uint8, limit uint8) []uint8 {
	numbers := make([]uint8, quantityNumbers)

	for i := 0; i < int(quantityNumbers); i++ {
		num := rand.Intn(int(limit)) + 1

		exists := false
		for j := 0; j < i; j++ {
			if numbers[j] == uint8(num) {
				exists = true
				break
			}
		}

		if exists {
			i--
		} else {
			numbers[i] = uint8(num)
		}
	}

	return numbers
}

package validations

import (
	"github.com/link2618/Go-lotery/models"
)

func IsValidBaloto(baloto models.Baloto) (string, bool) {
	if len(baloto.Type) == 0 {
		return "El tipo de baloto es obligatorio.", false
	}

	if baloto.Number1 == 0 || baloto.Number2 == 0 || baloto.Number3 == 0 || baloto.Number4 == 0 || baloto.Number5 == 0 || baloto.Serial == 0 {
		return "Los números no pueden ser 0 o null", false
	}

	if !IsInRange(baloto.Number1) || !IsInRange(baloto.Number2) || !IsInRange(baloto.Number3) || !IsInRange(baloto.Number4) || !IsInRange(baloto.Number5) {
		return "Los números no se encuentran en el rango de 1 a 44.", false
	}

	if !IsSerialInRange(baloto.Serial) {
		return "El serial no se encuentra en el rango de 1 a 16.", false
	}

	if !AreNumbersUnique(baloto.Number1, baloto.Number2, baloto.Number3, baloto.Number4, baloto.Number5) {
		return "No se pueden repetir los números.", false
	}

	return "", true
}

func IsInRange(number uint8) bool {
	return number >= 1 && number <= 44
}

func IsSerialInRange(serial uint8) bool {
	return serial >= 1 && serial <= 16
}

func AreNumbersUnique(numbers ...uint8) bool {
	seen := make(map[uint8]struct{})

	for _, num := range numbers {
		if _, ok := seen[num]; ok {
			return false
		}
		seen[num] = struct{}{}
	}

	return true
}

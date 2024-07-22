package ejercicios

import "strconv"

func ConvierteAEntero(texto string) (string, int) {
	entero, err := strconv.Atoi(texto)
	if err != nil {
		return "error", 0
	}
	
	if entero > 100 {
		return "Es mayor a 100", entero
	}

	return "Es menor a 100", entero
}
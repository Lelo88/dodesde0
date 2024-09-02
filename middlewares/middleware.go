package middlewares

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")	

	result := operacionesMidd(sumar)
	fmt.Println("Suma: ", result)

	result = operacionesMidd(restar)
	fmt.Println("Resta: ", result)

	result = operacionesMidd(multiplicar)
	fmt.Println("Multiplicar: ", result)
}

func operacionesMidd(f func(int, int) int) int {
	return f(2, 3)
}
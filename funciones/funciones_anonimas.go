package funciones

import "fmt"

//Calculo es una función anonima 
func Calculo() {

	var numero3 int = 10
	var numero4 int = 20

	// Ejemplo de una función anonima
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(2, 3))

	// Comportamiento de sobrecarga de funciones
	calculo = func(numero1, numero2 int) int {
		return numero1 * numero2 * numero3
	 }

	fmt.Println(calculo(6, 7))
}
package variables

import "fmt"

func MostrarEnteros() {
	// Se pueden declarar por asignación o por declaración
	var a int // Declaración
	var b = 2 // Asignación y declaración
	c := 3 // Asignación implicita

	intde32 := int32(4)
	intde64 := int64(5)

	// Imprimir
	fmt.Println(a, b, c)
	fmt.Println("entero de 32 = ", intde32)
	fmt.Println("entero de 64 = ", intde64)
}
package arreglos_slices

import "fmt"

var tabla [10]int  

var tabla2 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // tambien podemos hacerlo asignando de manera directa

func MuestroArreglos() {
	tabla[7] = 1 // asigno elementos
	tabla[2] = 43

	tablaString := [10]string{"uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve", "diez"}

	fmt.Println(tabla)
	fmt.Println(tabla2)
	fmt.Println(tablaString)

	for i := 0; i < len(tabla); i++ { // recorro el arreglo
		fmt.Println(tabla[i])
	}

	matriz := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // matriz de 3x3
	fmt.Println(matriz)
}
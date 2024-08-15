package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


// TablaDeMultiplicar devuelve el texto de la tabla de multiplicar
func TablaDeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	
	var numero int
	var err error
	var texto string // se lo agregamos para la explicación de la teoría de archivos

	for {
		fmt.Println("Ingrese el numero para la tabla de multiplicar: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue // si ingreso mal el dato, me pide el ingreso nuevamente hasta que el numero sea correcto
			} else {
				break // si el dato es correcto, salgo del bucle
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto = texto +  fmt.Sprintf("%d x %d = %d\n", i, numero, i*numero)
	} 

	return texto
}
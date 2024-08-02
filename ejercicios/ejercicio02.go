package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)



func TablaDeMultiplicar() {
	var numero int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

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
		fmt.Printf("%d x %d = %d\n", i, numero, i*numero)
	} 
}
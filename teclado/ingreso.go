package teclado

import (
	"bufio" // lectura de teclado y tratamiento de archivos
	"fmt"   // formato de escritura
	"os"    // lectura de archivos
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresarNumeros() {
	fmt.Println("Ingresar primer numero")
	scanner := bufio.NewScanner(os.Stdin) // lectura de teclado
	
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text()) // 
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	} // lectura de teclado
	
	fmt.Println("Ingresar primer numero")
	scanner = bufio.NewScanner(os.Stdin) // lectura de teclado
	
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text()) // 
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	} // lectura de teclado

	fmt.Println("Ingresar la leyenda: ")
	scanner = bufio.NewScanner(os.Stdin) // lectura de teclado
	
	if scanner.Scan() {
		leyenda = scanner.Text() //
		} // lectura de teclado

	fmt.Println(leyenda, numero1+numero2)
}

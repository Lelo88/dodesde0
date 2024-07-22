package main // si el archivo no se encuentra en una carpeta, el package es main

import (
	//"fmt" // importa el paquete fmt
	"fmt"

	"github.com/Lelo88/dodesde0/condicionales"
	"github.com/Lelo88/dodesde0/variables" // importa el paquete github.com/desde0/desde0
)

func main() { // funcion principal
	// fmt.Println("Hola mundo")

	variables.MostrarEnteros()

	variables.RestoVariables()

	estado, texto := variables.ConviertoATexto(2)
	fmt.Println("Estado: ", estado, " Texto: ", texto)

	estado, convertido := variables.ConviertoATexto(3)
	fmt.Println("Convertido: ", convertido)

	fmt.Println("Estado: ", estado)


	condicionales.FuncionamientoCondicional()

	condicionales.FuncionamientoSwitch()
}


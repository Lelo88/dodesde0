package variables

import (
	"fmt"
	"time"
)
// si declaramos variables afuera, deben ser con mayusculas

var Nombre string
var Estado bool
var Sueldo float64
var Fecha time.Time

func RestoVariables(){
	Nombre = "Jorge"
	Estado = true
	Sueldo = 2.5
	Fecha = time.Now()

	fmt.Println("nombre: ", Nombre)
	fmt.Println("estado: ", Estado)
	fmt.Println("sueldo: ", Sueldo)
	fmt.Println("fecha: ", Fecha)
}
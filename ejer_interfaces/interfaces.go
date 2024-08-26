package ejer_interfaces

import (
	"fmt"

	"github.com/Lelo88/dodesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un %s, y estoy respirando", hu.Sexo())
}
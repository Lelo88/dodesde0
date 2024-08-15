package funciones

import "fmt"

func tabla(valor int) func () int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia++ // incrementa la secuencia
		return numero * secuencia // multiplico el numero por la secuencia
	}
}

func LlamarClosure() {
	tablade1 := 2
	MiTabla := tabla(tablade1)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
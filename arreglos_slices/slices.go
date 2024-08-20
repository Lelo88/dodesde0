package arreglos_slices

import "fmt"

var tablaSlice []int = []int{1, 2, 3} // tambien podemos hacerlo asignando de manera directa
var arregloSlice[10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func MuestroSlice() {
	fmt.Println(tablaSlice)

	porcion := arregloSlice[2:4]
	fmt.Println(porcion)
	porcion2 := arregloSlice[:4]
	fmt.Println(porcion2)
	porcion3 := arregloSlice[2:]
	fmt.Println(porcion3)
}

func Capacidad() {
	// elementos es un slice de enteros con capacidad de 5 elementos. De esta manera creamos un slice
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d", len(elementos))
	fmt.Printf("Capacidad %d", cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d", len(nums))
	fmt.Printf("Capacidad %d", cap(nums)) // la capacidad se duplica cada vez que se llena, reservando más capacidad. 
	// Para el sistema operativo es mas costoso reasignar memoria para el slice.
}
package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func Iterar2() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
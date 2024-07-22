package condicionales

import (
	"fmt"
	"runtime"
)

func FuncionamientoCondicional() {
	if os := runtime.GOOS; os == "linux" {
		fmt.Println("Es linux")
	} else if os == "darwin" {
		fmt.Println("Es macos")
	} else {
		fmt.Println("Es windows")
	}
}
package condicionales

import (
	"fmt"
	"runtime"
)

func FuncionamientoSwitch() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Es linux")
	case "darwin":
		fmt.Println("Es macos")
	case "windows":
		fmt.Println("Es windows")
	default:
		fmt.Println("Otro sistema operativo")
	}
}
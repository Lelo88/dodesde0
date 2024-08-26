package deferpanic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	// El defer nos permite configurar cual va a ser la ultima instruccion a ejecutarse cuando salimos de una función
	fmt.Println("Este es el primer mensaje")
	defer fmt.Println("Este es el mensaje final")


	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic(){
	// El panic nos permite detener la ejecucion del programa
	a:=1	
	if a==1{
		panic("Error") 
	}
}

func EjemploRecover(){
	// El recover nos permite recuperar el panic
	a:=1
	defer func(){
		if r:=recover(); r!=nil{
			log.Fatal("Ocurrió un error que generó panic: ", r)
		}
	}()
	if a==1{
		panic("Se encontró el valor 1")
	}
	fmt.Println("Continua")
}
package main // si el archivo no se encuentra en una carpeta, el package es main

import "github.com/Lelo88/dodesde0/webserver"

// "github.com/Lelo88/dodesde0/mapas"
// "github.com/Lelo88/dodesde0/users"
//"github.com/Lelo88/dodesde0/deferpanic"
//"fmt"

// "github.com/Lelo88/dodesde0/goroutine"
//interfaces "github.com/Lelo88/dodesde0/ejer_interfaces"
//"github.com/Lelo88/dodesde0/modelos"

//"fmt" // importa el paquete fmt
// "fmt"

//"github.com/Lelo88/dodesde0/condicionales"
//"github.com/Lelo88/dodesde0/archivos"
//"github.com/Lelo88/dodesde0/arreglos_slices"
//"github.com/Lelo88/dodesde0/funciones"
//"github.com/Lelo88/dodesde0/ejercicios"
//"github.com/Lelo88/dodesde0/iteraciones"
//"github.com/Lelo88/dodesde0/teclado"
//"github.com/Lelo88/dodesde0/variables" // importa el paquete github.com/desde0/desde0

func main() { // funcion principal
	/* 	// fmt.Println("Hola mundo")

	   	variables.MostrarEnteros()

	   	variables.RestoVariables()

	   	estado, texto := variables.ConviertoATexto(2)
	   	fmt.Println("Estado: ", estado, " Texto: ", texto)

	   	estado, convertido := variables.ConviertoATexto(3)
	   	fmt.Println("Convertido: ", convertido)

	   	fmt.Println("Estado: ", estado)


	   	condicionales.FuncionamientoCondicional()

	   	condicionales.FuncionamientoSwitch()

	   	mensaje, numero := ejercicios.ConvierteAEntero("200")
	   	fmt.Println("El mensaje es: ", mensaje, " El numero es: ", numero) */

	// ejecutamos el paquete teclado

	// teclado.IngresarNumeros()

	// paquete iteraciones

	// ejercicios.TablaDeMultiplicar()

	//archivos.GrabaTabla()

	//archivos.SumaTabla()

	//archivos.LeoArchivo()

	// funciones.Calculo()

	//funciones.LlamarClosure()

	//funciones.Exponencia(12)

	// mapas.MuestroMapas()

	//users.AltaUsuario()

	//  Pedro := new(modelos.Hombre)
	// interfaces.HumanosRespirando(Pedro)

	// Maria := new(modelos.Mujer)
	//interfaces.HumanosRespirando(Maria)

	//deferpanic.VemosDefer()

	//deferpanic.EjemploPanic()
	//deferpanic.EjemploRecover()

	// canal1 := make(chan bool)

	// go goroutine.MiNombreLentooo("Leandro", canal1)
	// llamo a la rutina y se ejecuta por ser asíncrono, por lo que no sabemos si se ejecutó o no
	// Vamos a crear un pequeño código de ejemplo
	//defer func() {
		// no va a terminar la ejecución hasta que todos los 
	//	<-canal1 // el canal es el que está enviando información. Esto sería un await
	//}()
	
	// var x string
	// fmt.Scanln(&x) // el programa se detiene hasta que el usuario presione enter u otra tecla


	webserver.MiWebServer()
}

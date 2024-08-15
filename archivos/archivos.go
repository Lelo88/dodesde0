package archivos

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Lelo88/dodesde0/ejercicios"
)

var fileName string = "./archivos/txt/tabla.txt"

// GrabaTabla graba el archivo con el texto de la tabla de multiplicar
func GrabaTabla () {
	var texto string = ejercicios.TablaDeMultiplicar()

	archivo, err := os.Create(fileName) // aca grabo el archivo
	if err != nil {
		fmt.Println("No se pudo crear el archivo")
		return // salgo del método sin hacer nada
	}

	fmt.Fprintln(archivo, texto) // copio todo lo de texto en archivo
	archivo.Close() // cierro el archivo, ya que grabo el texto
}

// SumaTabla suma el archivo con el texto de la tabla de multiplicar
func SumaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()

	if !Append(fileName, texto) {
		fmt.Println("No se pudo crear el archivo")
		return
	}
}

// Append me devuelve true si se pudo escribir en el archivo
 func Append(fileName string, text string) bool {
	archivo, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644) // Openfile abre el archivo. El numero es el formato de permiso que se da en Linux
	if err  != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	_, err = archivo.WriteString(text)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	archivo.Close()
	return true
}

// LeoArchivo lee el archivo
func LeoArchivo() {
	file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    buf := new(bytes.Buffer)
    _, err = io.Copy(buf, file)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println(buf.String())
}
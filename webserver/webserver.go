package webserver

import "net/http"


func MiWebServer() {
	http.HandleFunc("/", home) // se llama a la funcion home
	http.ListenAndServe(":3000", nil) // se crea el servidor en el puerto 3000
}

// sintaxis basica para crear un webserver
func home(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./index.html") // esta en la misma carpeta que el paquete webserver
}
package mapas

import "fmt"

func MuestroMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["MX"] = "Mexico"
	paises["CO"] = "Colombia"
	paises["AR"] = "Argentina"
	paises["ES"] = "Espana"
	fmt.Println(paises)

	fmt.Println(paises["MX"]) // devuelve el valor de la llave MX

	campeonato := map[string]int{
		"Colombia": 7,
		"Mexico": 6,
		"Brasil":  3,
		"Peru": 2,
	}

	fmt.Println(campeonato) // no lo muestra ordenado como aparece en el mapa, sino que toma la clave y realiza un orden alfabetico

	for equipo, puntaje := range campeonato {
		fmt.Println(equipo +  " tiene un puntaje de: ", puntaje)
	}

	delete(campeonato, "Colombia")
	puntaje, ok := campeonato["Argentina"]
	fmt.Println("El puntaje capturado es: ", puntaje, " y el equipo existe? ", ok)
}
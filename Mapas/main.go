package main

import (
	"fmt"
	"maps"
)

func main() {

	mapa := make(map[string]int)
	mapa["Amin"] = 4
	mapa["espinoza"] = 8

	fmt.Println("mapa: ", mapa)

	version1 := mapa["espinoza"]
	fmt.Println("version 1: ", version1)

	version2 := mapa["Amin"]
	fmt.Println("version 2: ", version2)

	fmt.Println("longitud del mapa: ", len(mapa))

	_, dato := mapa["Amin"]
	fmt.Println("data: ", dato)

	delete(mapa, "amin")
	fmt.Println("mapa: ", mapa)

	clear(mapa)
	fmt.Println("mapa: ", mapa)

	nuevoMapa1 := map[string]int{"Amin": 4, "espinoza": 8}
	fmt.Println("nuevo mapa: ", nuevoMapa1)

	nuevoMapa2 := map[string]int{"Amin": 4, "espinoza": 8}
	fmt.Println("nuevo mapa: ", nuevoMapa2)

	if maps.Equal(nuevoMapa1, nuevoMapa2) {
		fmt.Println("los mapas son iguales")
	}

}

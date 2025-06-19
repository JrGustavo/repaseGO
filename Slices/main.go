package main

import (
	"fmt"
	"slices"
)

func main() {
	var arregloCadenas []string
	fmt.Println("datos:", "contenido", arregloCadenas, "condicion:", arregloCadenas == nil, "longitud:", len(arregloCadenas))

	// Inicializar el arreglo con 3 elementos
	arregloCadenas = make([]string, 3)
	arregloCadenas[0] = "a"
	arregloCadenas[1] = "b"
	arregloCadenas[2] = "c"

	fmt.Println("datos:", arregloCadenas)
	fmt.Println("contenido:", arregloCadenas)
	fmt.Println("condicion:", arregloCadenas == nil, "longitud:", len(arregloCadenas))

	// Agregar elementos al arreglo
	arregloCadenas = append(arregloCadenas, "d")
	arregloCadenas = append(arregloCadenas, "e")
	fmt.Println("valores finales:", arregloCadenas)
	fmt.Println("longitud:", len(arregloCadenas))

	// Comparar con otro arreglo
	segundoArreglo := []string{"a", "b", "c", "d", "e"}
	fmt.Println("segundo:", segundoArreglo)

	tercerArreglo := []string{"a", "b", "c", "d", "e"}

	if slices.Equal(segundoArreglo, tercerArreglo) {
		fmt.Println("Los arreglos son iguales.")
	} else {
		fmt.Println("Los arreglos son diferentes.")
	}
}

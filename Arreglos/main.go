package main

import "fmt"

func main() {
	var arreglo [5]int
	fmt.Println("Arreglo completo:", arreglo)

	arreglo[4] = 100
	fmt.Println("Arreglo completo:", arreglo)
	fmt.Println("Arreglo completo:", arreglo[4])

	fmt.Println("Tamaño del arreglo:", len(arreglo))

	lista := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lista completo:", lista)

	lista2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lista completo:", lista2)
	fmt.Println("Tamaño del arreglo:", len(lista2))

}

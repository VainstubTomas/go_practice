package main

import (
	"fmt"
)

func main() {
	fmt.Println("MATRICES")

	var array_de_prueba [10]int
	array_de_prueba[9] = 250

	//REFERENCIA AL ULTIMO VALOR
	fmt.Println(array_de_prueba[len(array_de_prueba)-1])
}

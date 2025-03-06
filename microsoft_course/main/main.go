package main

import (
	"fmt"
	"microsoft_course/paquetes_calc"
)

func main() {

	suma_interna_prueba := paquetes_calc.internalSum(8)
	fmt.Println(suma_interna_prueba)

	suma_de_prueba := paquetes_calc.Sum(2, 5)
	fmt.Println(suma_de_prueba)
}

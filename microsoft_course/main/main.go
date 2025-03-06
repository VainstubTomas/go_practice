package main

import (
	"fmt"
	"microsoft_course/paquetes_calc"
)

func main() {

	//en este caso no deja utilizar una funcion que es privada

	// suma_interna_prueba := paquetes_calc.internalSum(8)
	// fmt.Println(suma_interna_prueba)

	suma_de_prueba := paquetes_calc.Sum(2, 5)
	fmt.Println(suma_de_prueba)

	resta_de_prueba_A := paquetes_calc.Rest(2, 5)
	fmt.Println(resta_de_prueba_A)

	mult_de_prueba := paquetes_calc.Mult(2, 5)
	fmt.Println(mult_de_prueba)

	div_de_prueba := paquetes_calc.Div(5, 2)
	fmt.Println(div_de_prueba)
}

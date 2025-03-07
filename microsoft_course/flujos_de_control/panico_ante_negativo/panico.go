package main

import (
	"fmt"
)

func main() {
	fmt.Println("PANICO ANTE NUMERO NEGATIVO:")
	var numero_ingresado int
	fmt.Print("INGRESE UN NUMERO REAL:")
	fmt.Scanf("%d", &numero_ingresado)

	if numero_ingresado >= 0 {
		fmt.Printf("El valor que usted ingreso -> %d\n", numero_ingresado)
		fmt.Println("ES POSITIVO")
	} else if numero_ingresado < 0 {
		fmt.Printf("El valor que usted ingreso -> %d\n", numero_ingresado)
		fmt.Println("ES NEGATIVO - PANIC!")
		panic("El valor ingresado no es valido")
	}

}

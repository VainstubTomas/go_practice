package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Traductor de numeros romanos:")
	var numero_romano_ingresado string
	fmt.Println("ingrese un numero romano: ")
	fmt.Scanln(&numero_romano_ingresado)
	numero_romano_ingresado = strings.ToUpper(numero_romano_ingresado)
	fmt.Println(numero_romano_ingresado)

	valores_letras := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	fmt.Println(valores_letras)

	var resultado int

	for i := 0; i < len(numero_romano_ingresado); i++ {

		letra := string(numero_romano_ingresado[i])
		var valor int

		switch letra {
		case "M":
			valor = 1000
		case "D":
			valor = 500
		case "C":
			valor = 100
		case "L":
			valor = 50
		case "X":
			valor = 10
		case "V":
			valor = 5
		case "I":
			valor = 1
		}

		resultado += valor
	}

	fmt.Printf("Su numero romano equivale a: %v", resultado)
}

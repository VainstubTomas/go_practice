package main

import "fmt"

// El número 2 es el único primo par
func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Búsqueda de números primos:")
	fmt.Println("Prime numbers less than 20:")

	for number := 1; number <= 20; number++ {
		if findprimes(number) {
			fmt.Printf("%v ", number)
		}
	}
	fmt.Println() // Agrega un salto de línea al final
}

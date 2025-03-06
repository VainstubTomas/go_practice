package main

import (
	"fmt"
)

func main() {

	fmt.Println("SECUENCIA DE FIBONACCI HASTA EL NUMERO 30:")
	fmt.Println("")

	/*La secuencia dice que se parte del numero 1 y que el numero que le sigue es la suma del mismo con los 2 anteriores*/

	a := 0
	b := 1

	for i := 1; i <= 10; i++ {
		c := a + b
		a = b
		b = c
		fmt.Println(c)
	}
}

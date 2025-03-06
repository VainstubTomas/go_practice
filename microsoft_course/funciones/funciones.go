package main

import (
	"fmt"
	// "strconv"
)

// func funcion_sumar(number1 string, number2 string) int {
// 	int1, _ := strconv.Atoi(number1)
// 	int2, _ := strconv.Atoi(number2)
// 	return int1 + int2
// }

func operaciones_basicas(number1 int, number2 int)(sum int, rest int, mult int, div float64){
	sum = number1 + number2
	rest = number1 - number2
	mult = number1 * number2
	div = float64(number1) / float64(number2)
	return
}

func main() {
	fmt.Println(operaciones_basicas(1, 2))
}

package main

import (
	"fmt"
	"strconv"
)

func funcion_sumar(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return
}

func main() {
	fmt.Println(funcion_sumar("1", "2"))
}

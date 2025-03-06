package paquetes_calc

// var Version = "1.0"

func Sum(number1, number2 int) int {
	return number1 + number2
}

func Rest(number1, number2 int) (posible_resta_uno int) {
	posible_resta_uno = number1 - number2
	return
}

func Mult(number1, number2 int) int {
	return number1 * number2
}

func Div(number1, number2 int) float64 {
	/*
		num1 != 0 o num1 = 0
		num2 != 0 o num2 = 0

		Posibles divisiones:
			div_nula = 0 / 0 (si 1 y 2 == 0)
			si uno de los dos es cero, la division no se puede realizar
	*/

	if number1 == 0 && number2 == 0 {
		return 0
	} else if number1 == 0 && number2 != 0 {
		return 0
	} else if number2 == 0 {
		return 0
	}
	return float64(number1) / float64(number2)

}

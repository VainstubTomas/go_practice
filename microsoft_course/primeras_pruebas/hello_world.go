package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World")

	var numero int = 5
	fmt.Println(numero)

	i := 10
	fmt.Println(i)

	var suma int = numero + i
	const iva float64 = 1.21
	// var suma_con_iva float = suma * iva TIPO DE ERROR INCOMPATIBLE
	var suma_con_iva float64 = float64(suma) * iva
	fmt.Println(suma)
	//%f incluye el dato en la cadena ; .2 reduce la cantidad de decimales
	fmt.Printf("la suma con iva da un resultado de %.2f\n", suma_con_iva)

	fmt.Println(reflect.TypeOf(numero))

	var myBool bool = true
	fmt.Print(myBool)

	const mi_nombre string = "Tomás"
	fmt.Println(mi_nombre)

	// //condicional

	var numero_de_prueba int = 22

	if i == 10 && numero_de_prueba != 21 {
		fmt.Print(true)
	} else {
		fmt.Print(false)
	}

	// //array

	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	fmt.Println(array)

	var array2 [4]string
	fmt.Println(array2)

	// //map con estructura clave valor

	myMap := make(map[string]string)
	myMap["nombre"] = "Tomás"
	myMap["apellido"] = "Vainstub"
	myMap["edad"] = "21"
	fmt.Println(myMap)

	// //lista

	myList := list.New()
	myList.PushBack("front")
	myList.PushBack("midle")
	myList.PushBack("back")
	fmt.Println(myList.Front().Value)
	fmt.Println(myList.Back().Value)

	// //bucle

	for index := 0; index < len(array); index++ {
		fmt.Println(array[index])
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// //llamada a la funcion de ejemplo
	fmt.Println(ejemplo_funcion())

	// //estructuras - struct (vendrian siendo las clases)

	type Persona struct {
		name        string
		age         int
		nacionality string
	}

	persona1 := Persona{"Tomás", 21, "Argentina"}
	fmt.Println(persona1)

	persona2 := Persona{"Benjamin", 13, "Argentina"}
	fmt.Println(persona2)

}

//Ejemplo de funcion

func ejemplo_funcion() string {
	return "Funcion funcionando :)"
}

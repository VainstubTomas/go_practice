package main

import (
	"fmt"
)

func main() {
	fmt.Println("SEGMENTOS:")
	//segmento subyacente
	segmento_meses := []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}
	fmt.Println(segmento_meses)
	//segmentos
	segmento1 := segmento_meses[0:3]
	segmento2 := segmento_meses[3:6]
	segmento3 := segmento_meses[6:9]
	segmento4 := segmento_meses[9:12]
	fmt.Printf("LARGO: %v\n", len(segmento_meses))
	fmt.Printf("CAP: %v\n", cap(segmento_meses))
	fmt.Println("Segmentos a partir del subyacente")
	fmt.Println(segmento1)
	fmt.Println(segmento2)
	fmt.Println(segmento3)
	fmt.Println(segmento4)

	//anexo de nuevos elementos
	var anexos []int

	for i := 0; i < 10; i++ {
		anexos = append(anexos, i)
		fmt.Printf("AnexoN°%d: %v\n", i, anexos)
	}
}

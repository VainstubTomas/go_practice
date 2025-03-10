package main

import (
	"fmt"
	"log"
	"net/http"
)

// creacion de un tipo personalizado
type dollars float32

// metodo que hace uso del tipo personalizado
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// otro tipo personalizado con formato mapa
type database map[string]dollars

// metodo ServeHTTP con receptor database
// el metodo usa los datos del receptor para recorrerlos en bucle e imprimirlos
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// Implementacion de la interfaz en una API de servidor
func main() {
	//instancia de database inicializandola con nuevos valores (C-V)
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	/*Se inicia el servidor mediante http.ListenAndServe donde se define la direccion
	y el objeto db que implementa una version personalizaada del metodo ServeHTTP*/
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

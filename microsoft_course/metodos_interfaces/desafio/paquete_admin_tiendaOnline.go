package main

import (
	"fmt"
)

type Account struct {
	Name     string
	Lastname string
}

type Employee struct {
	Account
	credits float64
}

type account_admin interface {
	change_name(name string)
}

func (a *Account) change_name(name string) {
	a.Name = name
}

type credits_admin interface {
	add_credits(cant float64)
	remove_credits(cant float64)
	check_credits() float64
}

func (e *Employee) add_credits(cant float64) {
	e.credits += cant
}

func (e *Employee) remove_credits(cant float64) {
	if cant <= e.credits {
		e.credits -= cant
	}
}

func (e *Employee) check_credits() float64 {
	return e.credits
}

func main() {
	fmt.Println("Paquete para administrar una tienda en linea")

	a := Account{Name: "Tomás"}
	var admin_a account_admin = &a

	fmt.Printf("El nombre del usuario es %s\n, empleado de Empresa S.A.\n", a.Name)

	admin_a.change_name("Toto")

	fmt.Printf("Usted ha actualizado su nombre publico en la empresa a: %s\n", a.Name)

	e := Employee{Account: Account{Name: "Tomás"}, credits: 500}
	fmt.Printf("El nombre del empleado es %s y cuenta con %v creditos a favor.\n", e.Name, e.credits)

	//& lo que hace ese signo en la e es hacer de puntero haciendo referencia de que es un empleado en este caso
	var admin_e credits_admin = &e
	//ahora admin_e puede utilizar todos los metodos de employee
	admin_e.add_credits(500)
	fmt.Printf("(A) Ahora cuenta con %v creditos\n", e.credits)

	admin_e.remove_credits(500)
	fmt.Printf("(B) Ahora cuenta con %v creditos\n", e.credits)
}

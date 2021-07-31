package main

import (
	"fmt"
)

//Declarar funciones 1

func imprimirnombre(nombre string) {
	fmt.Println("Fuera de main")
	fmt.Println("El nombre es: ", nombre)
}

//Declarar funciones 2
func suma(n1 int, n2 int) int {
	return n1 + n2
}

//Declarar funciones 3
func resta(n1, n2 int) (r int) {
	r = n1 - n2
	return
}

func main() {
	imprimirnombre("Jose")
	fmt.Println("Dentro de main")

	fmt.Println(suma(9, 8))
	fmt.Println(resta(10, 5))

	//Firma de una funcion
	fmt.Printf("Funcion suma: %T\n", suma)
	fmt.Printf("Funcion resta: %T\n", resta)
}

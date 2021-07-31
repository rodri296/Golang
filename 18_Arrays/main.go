package main

import (
	"fmt"
)

func main() {
	//Arrays

	//Declarar Arrays

	var x [3]int
	fmt.Println(x)

	var k [3][2]float64
	fmt.Println(k)

	//
	//Asignar valor a un array
	x[1] = 25
	fmt.Println(x)

	// Acceder a u indice en concreto
	fmt.Println(x[1])

	// Declarar he inicializar Arrays 1
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	// Declarar he inicializar arrays 2
	j := [...]int{1, 2, 3, 4}
	fmt.Println(j)

	// Declarar he inicializar arrays 3
	i := [...]int{
		1,
		2,
		3,
		4,
	}
	fmt.Println(i)

	//Arrays de cualquier tipo
	f := [...]float64{1.2, 1.5, 8.3}
	fmt.Println(f)

	//Function Len() indica el tamaño del array
	fmt.Println(len(f))

	//Imprimir el ultimo elemento de un array
	fmt.Println(f[len(f)-1])

	// Comparar arrays
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 2, 4}
	//d := [4]int{1, 2, 4}

	fmt.Println(a == b)
	fmt.Println(a == c)
	//fmt.Println(a == d)
}

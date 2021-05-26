package main

import (
	"fmt"
)

func main() {

	//Operadores Aritmeticos
	// + Suma
	// - Resta
	// * Multiplicacion
	// % Resto
	a := 5
	b := 3

	fmt.Println("5 + 3 = ", a+b)
	fmt.Println("5 - 3 = ", a-b)
	fmt.Println("5 * 3 = ", a*b)
	fmt.Println("5 / 3 = ", a/b)
	fmt.Println("5 % 3 = ", a%b)

	c := 3

	fmt.Println("**************************************************************")
	fmt.Println("5 == 3 = ", a == b)
	fmt.Println("3 == 3 = ", b == c)
	fmt.Println("5 != 3 = ", a != b)
	fmt.Println("3 != 3 = ", c != b)
	fmt.Println("5 < 3 = ", a < b)
	fmt.Println("5 <= 3 = ", a <= b)
	fmt.Println("3 <= 3 = ", c <= b)
	fmt.Println("5 > 3 = ", a > b)

}

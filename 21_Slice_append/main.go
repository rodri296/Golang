package main

import (
	"fmt"
)

func main() {

	//append

	//Declaraciones el Slice x
	x := make([]byte, 4, 10)
	fmt.Println(x)

	// Inicializamos x
	x = []byte{'H', 'O', 'L', 'A'}
	fmt.Println(x)

	//Imprimimos x dandole formato UTF-8
	fmt.Printf("Slice x: %q \n", x)

	//Imprimimos cada elemento del slice x por separado
	//Utilizamos un bucle for
	for i := 0; i < len(x); i++ {
		fmt.Printf("Slice x[%d]: %q \n", i, x[i])
	}

	//Asignamos un espacio en blanco en el indice 5 del slice x
	//x[5] = ' ' //Error
	//fmt.Println(x)

	//Utilizamos append
	x = append(x, ' ')
	fmt.Println(x)

	//Imprimimos la capacidad de x
	fmt.Println("Cap(x)", cap(x))

	//Agregamos los byte "MUNDO" al slice x
	x = append(x, 'M', 'U', 'N', 'D', 'O')

	//Imprimimos x
	fmt.Printf("Slice x: %q \n", x)

	//Imprimimos la capacidad de x
	fmt.Println("Cap(x)", cap(x))

	//Imprimimos la longitud de x
	fmt.Println("len(x)", len(x))

	fmt.Println("******************************")

	//Declaramos el slice y
	var y []int

	//Creamos un bucle for
	for i := 1; i <= 15; i++ {
		//Agregamos el valor de i al slice y
		y = append(y, i)
		//Imprimimos el slice y
		fmt.Println("Slice y: ", y)
		//Imprimimos la longitud y la capacidad de y
		fmt.Printf("Len y: %d - Cap y: %d - Elementos: %d \n", len(y), cap(y), i)
	}
}

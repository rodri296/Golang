package main

import (
	"fmt"
)

//Closure

func print(cadena string) {
	fmt.Println(cadena)
}

func print2(cadena string) {
	fmt.Println(cadena)
}

func print3(cadena1, cadena2 string) {
	fmt.Println(cadena1 + cadena2)
}

func print4(fprint func(string)) {
	fprint("Hola mundo desde Print4")
}

func incrementar() func() int {
	i := 0
	return func() (r int) {
		r = i
		i += 2
		return
	}
}

func incremento() {
	i := 0
	i++
	fmt.Println(i)
}
func main() {
	cadena := "Hola Mundo"

	imprimir := print
	imprimir(cadena)

	imprimir2 := func() {
		fmt.Println(cadena)
	}
	imprimir2()

	imprimir = print2
	imprimir("Hola Mundo 2")

	//imprimir = print3

	fmt.Printf("Function print: %T\n", print)
	fmt.Printf("Function print: %T\n", print2)
	fmt.Printf("Function print: %T\n", print3)

	//print4(print)

	//Las funciones son comprables con nil
	var fb func()
	if fb == nil {
		fmt.Println("fb es igual a nil")
	}

	inc := incrementar()
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())

	inc2 := incremento()
	fmt.Println(inc2())
	fmt.Println(inc2())
}

package main

import (
	"fmt"
)

func main() {
	//Contador de números impares

	encabezado := `
	******************************
	Contador de numeros impares
	******************************
	`
	//Imprimimos el encabezado
	fmt.Println(encabezado)

	//Solicitamos el primer numero
	fmt.Println("Digita el primer numero")

	//Declaramos una variable para almacenar el numero digitado
	var numero1 int

	//Leemos el numero digitado y lo guardamos en la variable numero1
	fmt.Scanln(&numero1)

	fmt.Println("Digita el numero2")
	//Declaramos una variable para almacenar el numero digitado
	var numero2 int

	//Leemos el numero digitado y lo guardamos en la variable numero2
	fmt.Scanln(&numero2)

	//Declaramos la variable contador para guardar la cantidad de numeros primos
	var contador int

	//Realizamos un bucle tomando como inicio y fin los numeros digitados
	for i := numero1; i <= numero2; i++ {
		//Evaluamos si el numero es primo
		if i%2 != 0 {
			//Si el numero es primo
			//Incrementamos el valor variable contador en 1
			contador++

			//Imprimimos el numero primo
			fmt.Printf("%d es un numero impar\n", i)
		}
	}

	encabezado = `
	***************************
	RESULTADO DEL CONTEO
	***************************
	`

	//Imprimimos el encabezado
	fmt.Println(encabezado)

	//Imprimimos los resultados
	fmt.Printf("Entre el %d y el %d hay %d numeros impares.", numero1, numero2, contador)
}

package main

import (
	"fmt"
)

func main() {

	//Maps

	//Declarar Maps 1

	x := make(map[string]string)
	fmt.Println(x)

	//Declarar Maps 2
	y := make(map[string]string, 2)
	fmt.Println(y)

	//Agregar valores
	x["nombre"] = "Alejandro"
	x["edad"] = "29"
	fmt.Println(x["edad"])

	//Declarar Maps 3
	edades := map[string]int{
		"ana":       55,
		"rafael":    23,
		"manuel":    26,
		"alejandro": 29,
		"mario":     15,
	}
	fmt.Println(edades["ana"])

	//Para el indice podemos utilizar cualquier tipo que sea comparable con el operador ==
	//Para el valor podemos utilizar cualquier tipo

	dias := map[int]string{
		1: "Lunes",
		2: "Martes",
		3: "Miercoles",
	}
	fmt.Println(dias[2])

	//Eliminar elementos de un map
	delete(edades, "ana")
	fmt.Println(edades)

	//Si el elemento no existe los maps lo devuelven
	//El valor 0 del tipo de dato que sea el elemento
	fmt.Println("La edad de pedro es: ", edades["pedro"])

	//Podemos utilizar el operador ++
	edades["pedro"]++
	fmt.Println(edades)

	//saber si un elemento de un Map existe
	//Podemos utilizar los operadores de tipo == += -=
	if edad, ok := edades["manuel"]; ok {
		if edad < 18 {
			fmt.Println("Es menor de edad")
		} else {
			fmt.Println("Es mayor de edad")
		}
	} else {
		fmt.Println("El valor no existe")
	}

	//Saber el tamaño de un Map
	fmt.Println(len(edades))

	//Los tamaños de un map no son variables por lo que no podemos obtener sus direcciones
	//puntero := &edades["carlos"]
	//Recorrer un map
	//Los map son desordenados
	for nombre, edad := range edades {
		fmt.Printf("la edad de %s es: %d \n", nombre, edad)
	}
}

package main

import (
	"fmt"
)

func main() {

	//Estructuras de Control: Switch

	//Imprimir el nombre del dia de la semana que el usuaro digita

	// var dia int

	// fmt.Println("Digita el numero del dia de la semana")

	// fmt.Scanln(&dia)

	//	if dia == 1 {
	//		fmt.Println("Usted digitó Lunes")
	//	} else if dia == 2 {
	//		fmt.Println("Usted digitó el Martes")
	//	} else if dia == 3 {
	//		fmt.Println("Usted digitó el Miercoles")
	//	} else if dia == 4 {
	//		fmt.Println("Usted digitó el Jueves")
	//	} else if dia == 5 {
	//		fmt.Println("Usted digitó el Viernes")
	//	} else if dia == 6 {
	//		fmt.Println("Usted digitó el Sabado")
	//	} else if dia == 7 {
	//		fmt.Println("Usted digitó el Domingo")
	//	} else {
	//		fmt.Println("Digito un numero invalido")
	//	}

	//	fmt.Println("***************** Fin del Programa *******************")

	//switch dia {
	//case 1:
	//	fmt.Println("Usted digito Lunes")
	//case 2:
	//	fmt.Println("Usted digito Martes")
	//case 3:
	//	fmt.Println("Usted digito Miercoles")
	//case 4:
	//	fmt.Println("Usted digito jueves")
	//case 5:
	//	fmt.Println("Usted digito Viernes")
	//case 6:
	//	fmt.Println("Usted digito Sabado")
	//case 7:
	//	fmt.Println("Usted digito Domingo")
	//default:
	//	fmt.Println("Digito un numero invalido")
	//}

	//	fmt.Println("***************** Fin del Programa *******************")

	numero := 26

	switch {
	case numero > 23:
		fmt.Println("Mayor que 23")
		fallthrough
	case numero > 25:
		fmt.Println("Es mayor que 25")
	default:
		fmt.Println("Por lo menos es un Numero")
	}
}

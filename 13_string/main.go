package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cadena string
	cadena = "Hola Mundo"
	fmt.Println(cadena)

	fmt.Println(len(cadena) - 1)

	fmt.Println(cadena[9])

	fmt.Println(cadena[:])

	cadena = cadena + " nuevamente"

	fmt.Println(cadena)

	cadena += " siiiii"
	fmt.Println(cadena)

	cadena = `
	<html>
		<head>
			<meta charset="utf-8">
			<title></title>
		</head>
		<body>

		</body>
	</html>
	`
	fmt.Println(cadena)

	cadena = "Hola Mundo  \t\"25\""
	fmt.Println(cadena)

	edad := 29
	cadena = "La edad es " + strconv.Itoa(edad)
	fmt.Println("Edad", edad)
}

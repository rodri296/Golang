package main

import "fmt"

// hola mundo
/*
Hola Mundo
*/

func main() {
	var entero32 int32
	var entero64 int64

	entero32 = 25
	entero64 = 85

	fmt.Println(entero32 + int32(entero64))
}

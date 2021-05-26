package main

import (
	"fmt"
	"unsafe"
)

// hola mundo
/*
Hola Mundo
*/

func main() {
	var entero32 int32
	var entero64 int64
	//var enteroRune rune //alias para int32
	var enteroInt int //32 o 64 bits

	entero32 = 25
	entero64 = 85

	fmt.Println("07 tipos de datos")
	fmt.Println(entero32 + int32(entero64))

	//enteroRune = 46
	enteroInt = 56
	fmt.Println(entero32 + int32(enteroInt))
	//fmt.Fprintln(entero32 + enteroRune)

	fmt.Println(unsafe.Sizeof(enteroInt), unsafe.Sizeof(entero32))
}

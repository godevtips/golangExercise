package main

import "fmt"

func main() {

	fmt.Printf("%.[3]s valor de índice de argumento\n", 10)
	// resultado -> %!s(BADINDEX) valor de índice de argumento

	stringValue := "Go dev tips!"
	fmt.Printf("Hola %s\n", stringValue)
	// resultado -> Hola Go dev tips!

	valorHexadecimal := 26
	fmt.Printf("Hexadecimal minúsculo: %x\n", valorHexadecimal)  // resultado -> Hexadecimal minúsculo: 1a
	fmt.Printf("Hexadecimal mayúsculas: %X\n", valorHexadecimal) // resultado -> Hexadecimal mayúsculas: 1A

	valorHex := fmt.Sprintf("%x", valorHexadecimal)
	fmt.Println("Hexadecimal: ", valorHex) // resultado -> Hexadecimal:  1a
}

package main

import "fmt"

func main() {
	stringValue := "Go dev tips!"
	fmt.Printf("Hello %s\n", stringValue)

	hexaValue := 26
	fmt.Printf("Hexadecimal: %x\n", hexaValue)
	fmt.Printf("Hexadecimal 2: %X\n", hexaValue)

	//valueOnly := fmt.Sprintf("%x", hexaValue,hexaValue)
	valueOnly := fmt.Sprintf("%x", hexaValue)
	fmt.Println("Value only: ", valueOnly)
}

package main

import "fmt"

func main() {
	stringValue := "Go dev tips!"
	fmt.Printf("Hoi %s\n", stringValue) // Resultaat -> Hoi Go dev tips!

	hexWaarde := 26
	fmt.Printf("Hexadecimal in kleine letters: %x\n", hexWaarde) // Resultaat -> Hexadecimal in kleine letters: 1a
	fmt.Printf("Hexadecimal in hoofd letters: %X\n", hexWaarde)  // Resultaat -> Hexadecimal in hoofd letters: 1A

	hexadecimaalWaarde := fmt.Sprintf("%x", hexWaarde)
	fmt.Println("Hexadecimaal waarde: ", hexadecimaalWaarde) // Resultaat -> Hexadecimaal waarde:  1a
}

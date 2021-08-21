package main

import "fmt"

func main() {

	convertidaAUnicode := '\u0250' // -> convertida a unicode
	fmt.Printf("Convertida A unicode printf: %c\n", convertidaAUnicode)

	resultado := fmt.Sprintf("%c", convertidaAUnicode) // Si solo desea el resultado.
	fmt.Println("Convertida A unicode símbolo resultado: ", resultado)
}

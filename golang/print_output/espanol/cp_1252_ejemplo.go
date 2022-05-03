package main

import "fmt"

func main() {

	simbolo_euro := '\u20AC' // -> símbolo del euro
	fmt.Printf("Símbolo del euro printf: %c\n", simbolo_euro)

	resultado := fmt.Sprintf("%c", simbolo_euro) // Si solo desea el resultado.
	fmt.Println("Euro símbolo resultado: ", resultado)

}

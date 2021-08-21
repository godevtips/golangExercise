package main

import "fmt"

func main() {

	simbolo_euro := '\u20AC' // -> símbolo del euro
	fmt.Printf("Símbolo del euro printf: %q\n", simbolo_euro)

	resultado := fmt.Sprintf("%q", simbolo_euro) // Si solo desea el resultado.
	fmt.Println("Euro símbolo resultado: ", resultado)

}

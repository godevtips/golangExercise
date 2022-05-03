package main

import "fmt"

func main() {
	caracteres_decimal := 33
	fmt.Printf("Caracteres: %c\n", caracteres_decimal)            // !
	fmt.Printf("Caracteres citados: %q\n", caracteres_decimal)    // '!'
	fmt.Printf("Unicode: %U\n", caracteres_decimal)               // U+0021
	fmt.Printf("Unicode con carácter: %#U\n", caracteres_decimal) // U+0021 '!'

	valor := fmt.Sprintf("%U", caracteres_decimal) // Devuelve el valor resultante
	fmt.Println(valor)

}

package main

import "fmt"

func main() {
	character_decimal := 33
	fmt.Printf("Karakter: %c\n",character_decimal) // !
	fmt.Printf("Geciteerd karakter: %q\n",character_decimal) // '!'
	fmt.Printf("Unicode: %U\n",character_decimal) // U+0021
	fmt.Printf("Unicode met karakter: %#U\n",character_decimal) // U+0021 '!'
}

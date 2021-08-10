package main

import "fmt"

func main() {
	unicode_decimal := 33
	fmt.Printf("Character: %c\n",unicode_decimal) // !
	fmt.Printf("Quoted character: %q\n",unicode_decimal) // '!'
	fmt.Printf("Unicode: %U\n",unicode_decimal) // U+0021
	fmt.Printf("Unicode with character: %#U\n",unicode_decimal) // U+0021 '!'
}

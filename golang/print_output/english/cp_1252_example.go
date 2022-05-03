package main

import "fmt"

func main() {

	euroSign := '\u20AC' // -> euro sign unicode
	fmt.Printf("Euro sign print format: %c\n", euroSign)

	valueOnly := fmt.Sprintf("%c", euroSign) // If you only want the resulting string.
	fmt.Println("Euro sign value: ", valueOnly)
}

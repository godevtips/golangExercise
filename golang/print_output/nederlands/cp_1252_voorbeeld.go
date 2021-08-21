package main

import "fmt"

func main() {

	euroTeken := '\u20AC' // -> euro teken unicode
	fmt.Printf("Euro teken printf: %c\n", euroTeken)

	valueOnly := fmt.Sprintf("%c", euroTeken) // Om de resulterende string te retourneren.
	fmt.Println("Euro teken string: ", valueOnly)
}

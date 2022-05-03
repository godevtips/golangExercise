package main

import "fmt"

func main() {

	gedraaidAUnicode := '\u0250' // -> Latijnse kleine letter a gedraaid unicode
	fmt.Printf("Latijnse kleine letter a gedraaid simbool: %c\n", gedraaidAUnicode)

	gedraaidAUnicodeWaarde := fmt.Sprintf("%c", gedraaidAUnicode) // Om de resulterende string te retourneren.
	fmt.Println("Gedraaide A unicode waarde: ", gedraaidAUnicodeWaarde)

}

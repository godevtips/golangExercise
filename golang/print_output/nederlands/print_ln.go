package main

import "fmt"

func main() {

	fmt.Println("Dit is de", "eerste", "zin")
	fmt.Println("Dit is de", "tweede", "zin")
	// Output:
	// Dit is de eerste zin
	// Dit is de tweede zin

	fmt.Print("Dit is de ", "eerste ", "zin")
	fmt.Print("Dit is de ", "tweede ", "zin")
	// Output -> Dit is de eerste zinDit is de tweede zin
}

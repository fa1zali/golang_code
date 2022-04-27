package main

import "fmt"

func main() {

	// Declaring multiple variables on a single line
	var part1, part2 string

	part1 = "To be"
	part2 = "Or not to be"

	fmt.Println(part1 + " " + part2)

	// If we know the values to be assigned

	statement, isValid := "Penguins cannot fly.", true
	fmt.Println(statement)
	fmt.Println(isValid)
}

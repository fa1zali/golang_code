package main

import "fmt"

func main() {

	var basketTotal float64
	// normal addition
	var carrotPrice = 0.75
	basketTotal = basketTotal + carrotPrice
	fmt.Println(basketTotal)
	// Using shortcut operator
	var spinachPrice = 1.40
	basketTotal += spinachPrice
	fmt.Println(basketTotal)

	// Using shortcut operator with string
	command := "Hold my "
	beverage := "soda"
	command += beverage
	fmt.Println(command)
}

package main

import "fmt"

func main() {
	// Zero Type variables
	var classTime int32
	var averageGrade float32
	var teacherName string
	var isPass bool

	fmt.Println(classTime)
	fmt.Println(averageGrade)
	fmt.Println(teacherName)
	fmt.Println(isPass)

	// Inferring Type shorthand declaration operator :=
	nuclearMeltdownOccuring := true
	radiumInGroundWater := 4.521
	daySinceLastCatastrophe := 0
	externalMessage := "Everything is normal, Keep calm and carry on."

	fmt.Println(nuclearMeltdownOccuring)
	fmt.Println(radiumInGroundWater)
	fmt.Println(daySinceLastCatastrophe)
	fmt.Println(externalMessage)

	// Inferring Type
	var isItRaining = true
	var amountOfRainInCm = 2.25
	var maxDaysOfContinousRain = 2
	var season = "Monsoon"

	fmt.Println(isItRaining)
	fmt.Println(amountOfRainInCm)
	fmt.Println(maxDaysOfContinousRain)
	fmt.Println(season)
}

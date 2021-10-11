package main

import "fmt"

// You can init using conventional declare syntax
var initVar string = "Init"

// You CANNOT use short-hand inference syntax outside function scope
// initNewVar := "Init-again"

func main() {
	// Strings
	// var nameOne string = "MinhNPQ"
	// var nameTwo = "VNG"
	// var nameThree string

	// fmt.Println("Hello", nameOne, "from", nameTwo, nameThree)

	// nameThree = "Corporation"

	// fmt.Println("Hello", nameOne, "from", nameTwo, nameThree)

	// // Short-hand inference
	// nameFour := "Inferred Short-hand"

	// fmt.Println(nameFour)

	// Integers
	var ageOne int = 20
	var ageTwo int = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits & Memory => Not necessarily used besides int
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 256

	var scoreOne float32 = 6.9
	var scoreTwo float64 = 31293312312.3213 // Mostly USED
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}

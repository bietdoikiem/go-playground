package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Println("Hello,", name)
}

func sayBye(name string) {
	fmt.Println("Bye,", name)
}

func cycleNames(names []string, callback func(string)) {
	for _, value := range names {
		callback(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func main() {
	// var name string = "MinhNPQ"
	// sayGreeting(name)
	// sayBye(name)

	// listOfNames := []string{"MinhNPQ", "KhaBT", "Luigi", "Mario"}
	// cycleNames(listOfNames, sayGreeting)
	// cycleNames(listOfNames, sayBye)

	areaOne := circleArea(5)
	areaTwo := circleArea(15)
	fmt.Println("Area 1:", areaOne, "- Area 2:", areaTwo)
	fmt.Printf("Area of 1st circle is %v - Area of 2nd circle is %v", areaOne, areaTwo)
}

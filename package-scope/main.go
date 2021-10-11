package main

import "fmt"

var score float64 = 99.5

func main() {
	sayHello("Mario")

	showScore()

	for _, value := range points {
		fmt.Println(value)
	}
}

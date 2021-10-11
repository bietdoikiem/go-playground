package main

import "fmt"

var score = 99.5

func main() {
	sayHello("Mario")

	showScore()

	for _, value := range points {
		fmt.Println(value)
	}
}

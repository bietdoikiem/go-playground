package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	circumference() float64
}

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

// Square methods
func (s Square) area() float64 {
	return s.length * s.length
}

func (s Square) circumference() float64 {
	return s.length * 4
}

// Circle methods
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) circumference() float64 {
	return math.Pi * c.radius * 2
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumference())
}

func main() {
	shapes := []Shape{
		Square{length: 15.2},
		Circle{radius: 5.5},
		Square{length: 5.2},
		Circle{radius: 4.4},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}

package main

import "fmt"

func main() {
	// Arrays
	// var ages [3]int = [3]int{20, 18, 21}
	var ages = [3]int{20, 18, 21}

	names := [4]string{"David", "Minh", "Kha", "Kitty"}
	names[1] = "Luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood) without fixed length
	var scores = []int{100, 200, 300, 400}
	scores[1] = 500
	scores = append(scores, 700)

	fmt.Println(scores)

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Cooper")
	fmt.Println(rangeOne)
}

package main

import "fmt"

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is:", i)
	// }

	names := []string{"maria", "luigi", "henshin", "peach"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("The value at index %v is %v \n", index, value)
	// }

	// If you don't want to use the index
	// for _, value := range names {
	// 	fmt.Printf("The value is %v \n", value)
	// }

	// You cannot ALTER value in array while in loop as it's like a local copy!
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
		value = "dummyReplacement"
	}

}

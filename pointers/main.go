package main

import "fmt"

func updateName(name *string) {
	*name = "updated Minh"
}

func main() {
	name := "minh"

	// updateName(name)

	// fmt.Println("Memory address of name is", &name)

	m := &name // Declare a pointer that points to mem address of name variable

	fmt.Println("Memory address m is", m)

	fmt.Println("Value at memory address:", *m) // asterisk symbol (*) to dereference pointer to its value

	// Update name variable via pointer
	updateName(m)
	fmt.Println(name)
}

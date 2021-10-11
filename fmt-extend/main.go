package main

import "fmt"

func main() {
	age := 35
	name := "MinhNPQ"

	// // Print using escape newline
	// fmt.Print("Hello, MinhNPQ  \n")
	// fmt.Print("Hello VNG \n")

	// // Print with new line automatically
	// fmt.Println("Hello Engineering Platform")
	// fmt.Println("My name is", name, "with the age of", age)

	// Printf (Formatted string) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("Age is of type %T \n", age)
	fmt.Printf("You scored %0.1f points! \n", 222.255)

	// Sprintf (save formatted string)
	var str string = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println(str)
}

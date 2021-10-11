package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("Minh Nguyen")
	fmt.Printf("First Name: %v - Last name: %v \n", fn1, sn1)

	fn2, sn2 := getInitials("Kha Bui")
	fmt.Printf("First Name: %v - Last name: %v \n", fn2, sn2)

	fn3, sn3 := getInitials("Cloud")
	fmt.Printf("First Name: %v - Last name: %v \n", fn3, sn3)
}

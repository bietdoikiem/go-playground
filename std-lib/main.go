package main

import (
	"fmt"
	"sort"
)

func main() {
	// "strings" library - utilities for string type
	// greeting := "Hello there, friends!"
	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "friends", "MinhNPQ"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th"))
	// fmt.Println(strings.Split(greeting, ","))

	// Original string
	// fmt.Println("Original string =", greeting)

	ages := []int{45, 20, 12, 60, 321, 59, 2}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 18)
	fmt.Println(index)

	names := []string{"minh", "kha", "pham", "nguyen"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "kha"))
}

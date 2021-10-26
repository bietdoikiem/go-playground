package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"coffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Loop through maps
	for k, v := range menu {
		fmt.Printf("Item %s has value of %f \n", k, v)
	}

	// Ints as key type
	phonebook := map[int]string{
		934717924: "MinhNPQ",
		279993724: "KhaBui",
	}

	fmt.Println(phonebook[934717924])

	// Mutate map value
	phonebook[934717924] = "Minh Pro"
	fmt.Println(phonebook[934717924])
}

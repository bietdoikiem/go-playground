package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 40)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is above or equal 40")
	}

	names := []string{"Mario", "Henshin", "Luigi", "MinhNguyen", "KhaBui"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}
		fmt.Printf("The value at pos %v is %v \n", index, value)
	}

}

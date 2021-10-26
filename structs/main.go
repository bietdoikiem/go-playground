package main

import "fmt"

func main() {
	myBill := newBill("Minh's Bill")

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("chicken rice", 8.50)
	myBill.updateTip(10)

	fmt.Println(myBill.format())
}

package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Format the bill
func (b *Bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	// Tip
	fs += fmt.Sprintf("%-25v ... $%v \n", "tip:", b.tip)

	// Total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total+b.tip)
	return fs
}

// Update tip
func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

// Add an item to the bill
func (b *Bill) addItem(item string, price float64) {
	b.items[item] = price
}

// Save bill
func (b *Bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}

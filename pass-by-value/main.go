package main

import (
	"fmt"
)

func updateName(name string) string {
	updatedName := fmt.Sprintf("Updated %s", name)
	return updatedName
}

func updateMenu(updatedMenu map[string]float64) {
	updatedMenu["coffee"] = 4.99
}

func main() {
	// Group A types -> strings, ints, booleans, floats, arrays, struct(s)
	name := "minh"

	name = updateName(name)
	fmt.Print("Name: ")
	fmt.Println(name)

	// Group B types -> slices, maps, functions
	// 1. Maps
	menu := map[string]float64{
		"pie":       6.44,
		"ice cream": 2.44,
	}

	updateMenu(menu)
	fmt.Print("Menu map: ")
	fmt.Println(menu)

	// 2. Slices
	var fieldPlayers int = 11
	var allPlayers int = 22
	footballTeam := make([]string, fieldPlayers, allPlayers)
	fmt.Printf("Football Team players on field: %d - Total players including substitution: %d\n", len(footballTeam), cap(footballTeam))

	// Loop and add on-field players
	for i := 0; i < int(fieldPlayers); i++ {
		fieldPlayer := fmt.Sprintf("FieldPlayer #%d", i+1)
		footballTeam[i] = fieldPlayer
	}
	// Loop and continue to add total players including subtitions to the team
	var continueIndex int = cap(footballTeam) - len(footballTeam)
	for i := continueIndex; i < cap(footballTeam); i++ {
		subPlayer := fmt.Sprintf("SubPlayer #%d", i+1)
		footballTeam = append(footballTeam, subPlayer) // Use append to expand the length to the given capacity
	}

	fmt.Println(footballTeam)
}

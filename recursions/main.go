package main

import "fmt"

var tick int = 0

const Limit = 10

func recursiveTickLogs(currentTick *int) int {
	fmt.Println("Current tick is", *currentTick)
	*currentTick++
	if *currentTick == Limit {
		fmt.Println("Last limit tick is", *currentTick)
		return *currentTick
	}
	return recursiveTickLogs(currentTick)
}

func main() {
	recursiveTickLogs(&tick)
}

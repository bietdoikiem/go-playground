package main

import "fmt"

// func count(thing string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}

// 	close(c)
// }

// Worker function for Worker Pool pattern
// To reduce bugs in channel communicate as we're processing and sending the results via different channels
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// I. Add wait group for GoCoroutines
	// var wg sync.WaitGroup
	// wg.Add(1)

	// // Go coroutines
	// go func() {
	// 	count("Sheep")
	// 	wg.Done()
	// }()
	// wg.Wait()

	// II. Channel (Blocking/Un-buffered)
	// c := make(chan string)
	// go count("Sheep", c)

	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	// II. Channel (Non-Blocking/Buffered)
	// c := make(chan string, 2) // Channel with unfilled capacity (dummy corresponding receiver) to avoid blocking until it's full
	// c <- "Hello"
	// c <- "World"

	// msg := <-c
	// fmt.Println(msg)

	// msg = <-c
	// fmt.Println(msg)

	// III. Select statement
	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	for {
	// 		c1 <- "Every 500s"
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		c2 <- "Every two seconds"
	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()

	// // Select statements help picks out faster co-routine; makes it not dependent on each other
	// for {
	// 	select {
	// 	case msg := <-c1:
	// 		fmt.Println(msg)
	// 	case msg := <-c2:
	// 		fmt.Println(msg)
	// 	}
	// }

	// IV. Worker Pool pattern
	jobs := make(chan int, 50)
	results := make(chan int, 50)

	// Utilize concurrent GoRoutines
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 50; i++ {
		jobs <- i
	}
	// close(jobs)

	for j := 0; j < 50; j++ {
		fmt.Println(<-results)
	}
}

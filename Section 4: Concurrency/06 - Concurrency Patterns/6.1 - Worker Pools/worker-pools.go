package main

import "fmt"

func main() {
	jobs := make(chan int, 45)
	results := make(chan int, 45)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 45; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 45; i++ {
		fmt.Println(<-results)
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		write("Hello World")
		waitGroup.Done()
	}()

	go func() {
		write("Hello World 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()

	/*go write("Hello World")
	write("Programing in Go")*/
}

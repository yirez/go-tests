package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	fmt.Printf("cpu: %v\n", runtime.NumCPU())
	fmt.Printf("go routines: %v\n", runtime.NumGoroutine())

	waitGroup.Add(1)
	go shout() //concurrency
	yell()
	fmt.Printf("go routines: %v\n", runtime.NumGoroutine())

	waitGroup.Wait()
}

func shout() {
	for i := 0; i < 15; i++ {
		fmt.Printf("shouting Hi, %v times\n", i)
	}
	waitGroup.Done()
}
func yell() {
	for i := 0; i < 15; i++ {
		fmt.Printf("YELLING HI, %v times\n", i)
	}
}

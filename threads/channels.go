package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitGroup2 sync.WaitGroup
var channelInt = make(chan int)

func main() {

	waitGroup2.Add(2)
	go generateNum()
	go yell2()
	fmt.Printf("rand: %v \n", <-channelInt)
	close(channelInt)
	waitGroup2.Wait()
	fmt.Printf("%v go routines finished\n", runtime.NumGoroutine())
}

func generateNum() {
	defer waitGroup2.Done()
	rand.Seed(time.Now().UnixNano())
	channelInt <- rand.Intn(100-10) + 10
}

func yell2() {
	defer waitGroup2.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("YELLING HI, %v times\n", i)
	}
}

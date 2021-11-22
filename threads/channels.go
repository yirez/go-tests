package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup2 sync.WaitGroup

var channelInt = make(chan int)

//var channelInt = make(<-chan int)   //receive only
//var channelInt = make(chan<- int)   //send only

func main() {

	waitGroup2.Add(2)
	go generateNum()
	go yell2()
	fmt.Printf("rand: %v \n", <-channelInt)

	waitGroup2.Wait()
	go stuffToChannel()

	for val := range channelInt {
		fmt.Printf("in range reading channel:%v\n", val)
	}
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

func stuffToChannel() {
	for i := 0; i < 10; i++ {
		channelInt <- i
		fmt.Printf("stuffing %v to channel\n ", i)
		time.Sleep(1 * time.Second)
	}
	close(channelInt)
}

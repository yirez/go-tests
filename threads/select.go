package main

import (
	"fmt"
	"time"
)

func main() {

	var channelNumInt = make(chan int)
	var channelQuit = make(chan int)
	go stuffToChannelSelect(channelNumInt, channelQuit)

	for {
		select {
		case data := <-channelNumInt:
			fmt.Printf("in range reading channel:%v\n", data)
		case data := <-channelQuit:
			fmt.Printf("Exiting:%v\n", data)
			return
		}
	}
}

func stuffToChannelSelect(channelNumInt chan int, channelQuit chan int) {
	for i := 0; i < 10; i++ {
		channelNumInt <- i
		fmt.Printf("stuffing %v to channel\n ", i)
		time.Sleep(1 * time.Second)
	}

	channelQuit <- 0
}

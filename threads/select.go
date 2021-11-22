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
		case data, ready := <-channelNumInt:
			fmt.Printf("in range reading channel:%v - channel ready?%v\n", data, ready)
		case data, ready := <-channelQuit:
			fmt.Printf("Exiting:%v- channel ready?%v\n", data, ready)
			return
		}
		//ready will be false if channel closed or not ready yet for a reason
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

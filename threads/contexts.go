package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func main() {
	contextCancel, cancel := context.WithCancel(context.Background())
	contextCancel = context.WithValue(contextCancel, "myData", 42)
	log.Println("****\nprocessing cancelly request")
	go workThatIsCancelled(contextCancel)

	time.Sleep(time.Second * 3)
	log.Println("cancelling context")
	cancel()

	//this is needed so main doesn't destroy itself until
	//we see the ctx done message
	time.Sleep(time.Second * 1)

	log.Println("****\nprocessing deadline request")
	contextDeadline, cancelDeadline := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	go deadlineBreakingWork(contextDeadline)
	time.Sleep(time.Second * 4)
	if errors.Is(contextDeadline.Err(), context.DeadlineExceeded) {
		log.Println("Deadline exceeded, context cancelled at deadline")
	}
	cancelDeadline()
}

func workThatIsCancelled(contextCancel context.Context) {
	for i := 0; i >= 0; i++ {
		select {
		default:
			log.Printf("processing index:%v\n", i)
			time.Sleep(1 * time.Second)
			if i == 10 { //finish context completion at an arbitrary logic
				contextCancel.Done()
			}
		case <-contextCancel.Done():
			log.Printf("request cancelled. Value at context myData:%v\n", contextCancel.Value("myData"))
			return
		}
	}
}

func deadlineBreakingWork(contextDeadline context.Context) {
	log.Println("doing lots of work")
	time.Sleep(time.Second * 3) //just over the deadline of 2 secs so we drop to timeout
	contextDeadline.Done()      //normally you would exit stating you are done
}

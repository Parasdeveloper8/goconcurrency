package main

import (
	"fmt"
	"time"
)

// This goroutine is running in background consuming power,memory which is example of
// goroutine leak
// We have mechanism to cancel children goroutine called done channel
func doWork(done <-chan bool) {
	for {
		select {
		case <-done: //If it receives message from done channel
			return
		default:
			fmt.Println("Doing work")
		}
	}
}
func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 5)
	//close channel
	close(done)
	/*go func() {
		for {
			fmt.Println("Doing work")
		}
	}()
	time.Sleep(time.Hour * 100)*/
}

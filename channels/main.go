package main

import "fmt"

//Channels are used to communicate among go routines as they are independent
func main() {
	//let's create a channel
	chn := make(chan string)

	go func() {
		chn <- "data"
	}() //As it is a anonymous function we have to put bracket to invoke this
	msg1 := <-chn
	fmt.Println(msg1)

	go func() {
		chn <- "peta"
	}()
	msg := <-chn
	fmt.Println(msg)
}

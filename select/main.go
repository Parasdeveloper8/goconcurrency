package main

import "fmt"

//We will communicate with multiple channels ,so we use select statement
func main() {
	//channel one
	chnfirst := make(chan string)

	//channel two
	chntwo := make(chan string)

	go func() {
		chnfirst <- "first"
	}()

	go func() {
		chntwo <- "second"
	}()

	select {
	//if message comes from chnfirst
	case chnf := <-chnfirst:
		fmt.Println(chnf)
		//if message comes from chntwo
	case chnt := <-chntwo:
		fmt.Println(chnt)
	}
}

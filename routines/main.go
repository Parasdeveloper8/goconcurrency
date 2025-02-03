package main

import (
	"fmt"
	"time"
)

func sums(a, b int) {
	fmt.Println(a + b)
}

// Go follows fork-join model in which child is forked out and joined later
func main() {
	go sums(2, 4)
	go sums(3, 4)
	go sums(4, 4)
	time.Sleep(time.Second * 5)
	fmt.Println("Hello")
}

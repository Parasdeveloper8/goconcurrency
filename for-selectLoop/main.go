package main

import "fmt"

func main() {
	//we are using buffered channel here(limited capacity)
	charChn := make(chan string, 4)
	sliceAr := []string{"a", "b", "c"}
	//We will loop over slice usinf for loop

	for _, v := range sliceAr {
		select {
		case charChn <- v:
		}
	}
	close(charChn)

	for result := range charChn {
		fmt.Println(result)
	}
}

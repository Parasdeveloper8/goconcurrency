package main

import "fmt"

func slicetoChannel(nums []int) <-chan int {
	numchan := make(chan int)
	go func() {
		for _, n := range nums {
			numchan <- n
		}
		close(numchan)
	}()
	return numchan
}

func sqChn(data <-chan int) <-chan int {
	numchan := make(chan int)
	go func() {
		for n := range data {
			numchan <- n * n
		}
		close(numchan)
	}()
	return numchan
}

func main() {
	sliceAr := []int{1, 2, 3, 4, 5}
	//Stage 1
	firstChannel := slicetoChannel(sliceAr)
	//stage 2
	secondChannel := sqChn(firstChannel)
	//stage 3
	for n := range secondChannel {
		fmt.Println(n)
	}

}

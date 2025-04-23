package main

import (
	"fmt"
	"sync"
)

func createDB(wg *sync.WaitGroup) {
	defer wg.Done() //states that work has done
	fmt.Println("Task 1 DB created")
}

func startDB(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 2 DB started")
}

func queryDB(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 3 DB queried")
}

func closeDB(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task 4 DB closed")
}

func main() {
	var wg sync.WaitGroup //It is a type of checklist for goroutines

	wg.Add(4)

	go createDB(&wg)
	go startDB(&wg)
	go queryDB(&wg)
	go closeDB(&wg)

	wg.Wait() //wait till finish
	fmt.Println("Main function finished")
}

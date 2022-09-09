package main

import (
	"fmt"
	"time"
)

func p(x int) {
	// Create new variables
	count := 0
	now := time.Now()

	// Create and Initialize array
	start := now
	arr := []int{}
	for i := 0; i < x+1; i++ {
		arr = append(arr, i)
		count++
	}
	fmt.Println(arr)
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	// Create and Initialize Slice
	start = now
	slice := make([]int, x+1, x+1)
	for i := 0; i < x+1; i++ {
		slice[i] = i
		count++
	}
	fmt.Println(slice)
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	// Create and Initialize map
	start = now
	mymap := make(map[int]int)
	for i := 0; i < x+1; i++ {
		mymap[i] = i
		count++
	}
	fmt.Println(mymap)
	elapsed = time.Since(start)
	fmt.Println(elapsed)
}

func main() {
	// Take user input
	fmt.Printf("Enter an integer: ")
	var x int
	fmt.Scanln(&x)

	// Call function p
	p(x)
}

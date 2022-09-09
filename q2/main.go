package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func p1(x int) {
	// Create new variables
	var sumSlice int = 0
	now := time.Now()

	// Declare WaitGroup for Parallelism
	// var wg sync.WaitGroup

	// Start Time
	start := now

	// Create and Initialize Slice
	slice := make([]int, x, x)
	for i := 0; i < x; i++ {
		// Generate x random integers and store them in the Slice
		// wg.Add(1)
		slice[i] = rand.Intn(200)
		sumSlice = sumSlice + slice[i]
	}
	// Print Sum
	fmt.Println("Original slice:", slice)
	fmt.Println("Total Sum: ", sumSlice)

	// Output Time
	elapsed := time.Since(start)
	fmt.Println("Time: ", elapsed)

	// Parallel approach
	start = now
	totalSum := 0
	sum1 := 0
	sum2 := 0
	for i := 0; i < x; i++ {
		if i == x/2 {
			sum1 += slice[i]
		} else {
			sum2 += slice[i]
		}
		totalSum = sum1 + sum2
	}
	fmt.Println("Second Summation: ")
	fmt.Println("Total Sum: ", totalSum)
	elapsed = time.Since(start)
	fmt.Println("Time: ", elapsed)
	// wg.Wait()
}

func p2(x int) {
	// Create new variables
	now := time.Now()

	// Create and Initialize slice
	slice := make([]int, x, x)
	for i := 0; i < x; i++ {
		// Generate x random integers and store them in the Slice
		slice[i] = rand.Intn(200)
	}
	fmt.Println("Original Slice: ", slice)

	// Start Time
	start := now

	// Sort: Slice
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	fmt.Println("Slice: ", slice)

	// Output Time
	elapsed := time.Since(start)
	fmt.Println("Time: ", elapsed)

	// Sort: SliceStable
	start = now
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	fmt.Println("SliceStable: ", slice)
	elapsed = time.Since(start)
	fmt.Println("Time: ", elapsed)
}

func main() {
	// Take user input
	fmt.Printf("Enter an integer: ")
	var x int
	fmt.Scanln(&x)

	// Call function p
	p1(x)
	fmt.Println("\nMoving on to Part 2")
	fmt.Println("---------------------------------------------")
	p2(x)
	fmt.Println("Program executed successfully")
}

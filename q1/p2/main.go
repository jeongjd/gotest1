package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	fmt.Println("Enter an integer: ")
	var x int
	fmt.Scanln(&x)

	arr := []int{}
	var i int
	for i = 0; i < x+1; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)

	slice := make([]int, x+1, x+1)
	for i = 0; i < x+1; i++ {
		slice[i] = rand.Intn(100)
	}
	fmt.Println(slice)
}
Footer

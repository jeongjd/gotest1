package main

import (
	"fmt"
	// import "time"
)

func main() {
	fmt.Println("Enter an integer: ")
	var x int
	fmt.Scanln(&x)

	arr := []int{}
	var i int
	for i = 0; i < x+1; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)

	slice := make([]int, x+1, x+1)
	for i = 0; i < x+1; i++ {
		slice[i] = i
	}
	fmt.Println(slice)

	var mymap = make(map[int]int)
	for i = 0; i < x+1; i++ {
		mymap[i] = i
	}
	fmt.Println(mymap)
}

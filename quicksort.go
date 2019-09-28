package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fillArray(arr []int) {
	max := 500
	min := 1
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(max-min) + min
	}
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print("\t")
	}
}

/*
Write a program that randomly
generate a one-dimensional array, and implements
quicksort using recursion. Display the original
array, and the new array.
*/
func program1() {
	// Set up seed so it's different every time
	rand.Seed(time.Now().UnixNano())
	// Make the 1d array to sort
	size := 20
	arr := make([]int, size)
	fillArray(arr)
	printArray(arr)
	quicksort(arr)
	fmt.Print("\nAFTER SORT\n")
	printArray(arr)
}

// Recursive Quicksort
func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// set up left pointer
	left := 0
	// set up right pointer
	right := len(arr) - 1
	// choose a random index in the array
	point := rand.Intn(len(arr))
	// swap your point and rightmost values
	arr[point], arr[right] = arr[right], arr[point]
	// loop through the array
	for i, _ := range arr {
		if arr[i] < arr[right] {
			// swap the values
			arr[left], arr[i] = arr[i], arr[left]
			// move left over one spot
			left++
		}
	}
	// swap your pointer indecies
	arr[left], arr[right] = arr[right], arr[left]
	// now sort the slice of everything before the left
	quicksort(arr[:left])
	// now sort the slice of everything from left+1 after
	quicksort(arr[left+1:])

	return arr
}

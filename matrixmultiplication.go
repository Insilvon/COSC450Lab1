package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Prints the content of the 3x3 matrix
func printMatrix(m [][]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(m[i][j])
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

// Fills the 3x3 array with random numbers
func fillMatrix(m [][]int) {
	max := 5
	min := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = rand.Intn(max-min) + min
		}
	}
}

// Multiplies m1 x m2, filling results in m3
func multiplyMatrix(m1 [][]int, m2 [][]int, m3 [][]int) {
	// for each row in m1
	for i := 0; i < 3; i++ {
		// for each column in m2
		for j := 0; j < 3; j++ {
			// for each column item in m1
			for k := 0; k < 3; k++ {
				m3[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
}

// Main method - generates 3 2d arrays of size 3x3 and multiplies them together.
func program2() {
	// Set up seed so it's different every time
	rand.Seed(time.Now().UnixNano())
	// Declare m1
	m1 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		m1[i] = make([]int, 3)
	}
	fillMatrix(m1)
	fmt.Println("==================")
	fmt.Println("=====Matrix 1=====")
	fmt.Println("==================")
	printMatrix(m1)

	// Declare m2
	m2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		m2[i] = make([]int, 3)
	}
	fillMatrix(m2)
	fmt.Println("==================")
	fmt.Println("=====Matrix 2=====")
	fmt.Println("==================")
	printMatrix(m2)

	// Declare and fill m3
	m3 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		m3[i] = make([]int, 3)
	}

	multiplyMatrix(m1, m2, m3)
	fmt.Println("==================")
	fmt.Println("=====Matrix 3=====")
	fmt.Println("==================")
	printMatrix(m3)
}

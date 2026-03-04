package main

import "fmt"

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers.
func Subtract(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Go Basics Examples")
	fmt.Println("==================")

	// Variables
	name := "Go"
	version := 1.22
	fmt.Printf("Welcome to %s %.2f!\n", name, version)

	// Functions
	fmt.Printf("Add(2, 3) = %d\n", Add(2, 3))
	fmt.Printf("Subtract(10, 4) = %d\n", Subtract(10, 4))
}

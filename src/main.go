package main

import "fmt"

func main() {
	fmt.Println("Hello Sanjay!! Welcome to Jenkins Pipeline")
	addval := Add(4, 6)
	fmt.Println("Example for Addition ==================> 4 + 6 = ", addval)
	subval := Subtract(6, 4)
	fmt.Println("Example for Subtraction ==================> 6 - 4 = ", subval)
}

// Add is our function that sums two integers
func Add(x, y int) int {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) int {
	return x - y
}

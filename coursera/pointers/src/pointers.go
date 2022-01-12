package main

import "fmt"

func learnPointerAssignment() {
	var x int = 1
	var y int

	fmt.Printf("\nInitial Values\n")
	fmt.Printf("x = %d | y = %d \n", x, y)

	var ip *int = &x // ip is pointer to int with the address of x
	y = *ip          // y is assigned value in the address ip (which is 1)
	fmt.Printf("\nAfter Pointer Assignment\n")
	fmt.Printf("x = %d | y = %d\n", x, y)
}

func main() {
	learnPointerAssignment()
}

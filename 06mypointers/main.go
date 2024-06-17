package main

import "fmt"

func main() {
	fmt.Println("Welcome to understand pointers")
	// resposible for holiding interger into it
	var ptr *int
	fmt.Println("Value of the pointer is:   ", ptr)
	myNumber := 23
	var ptr2 *int = &myNumber
	fmt.Println("Value of the pointer is: ", ptr2)
	fmt.Println("Value to which the pointer points is : ", *ptr2)
	fmt.Println("Value to which the pointer points is : ", &ptr2)

}

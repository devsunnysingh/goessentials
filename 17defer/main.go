package main

import (
	"fmt"
)

// NOTES
// LIFO Last In First Out
// Defer is way of deferring any method and they work on LIFO mechanism

func main() {
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("Two")

	fmt.Println("hello")

	myDefer()
	// Explaination of defer working
	// hello is printed
	// myDefer() is called
	// all the printing for loop print is sent to defer
	// Since the method was called, now the method will complete. Thus, all the numbers will be printed
	// Two is printed as it was last inserted before the method was invoked
	// One is printed
	// World is printed

}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Value of i is ", i)

	}
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to function demp")
	greeter()
	result := adder(2, 3)
	fmt.Printf("Result is %v", result)

	proAddition := proAdder(1, 2, 3, 4, 5)
	fmt.Printf("Result is %v", proAddition)

}
func greeter() {
	fmt.Println("Hello from Greeter function")
}
func greeterTwo() {
	fmt.Println("This is method 2")
}
func adder(v1 int, v2 int) int {
	sum := v1 + v2
	return sum
}

// pro functions

func proAdder(values ...int) int {
	total := 0
	for i, _ := range values {
		total += values[i]
	}
	return total

}

package main

import "fmt"

func main() {
	fmt.Println("If else demo")
	loginCount := 22
	var result string
	if loginCount < 10 {
		result = "REgular user"
	} else {
		result = "NOt a regular users"
	}
	fmt.Println(result)

}

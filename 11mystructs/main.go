package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")
	sunny := User{"Sunny", "devsunny@dev.ca", true, 25}
	fmt.Println(sunny)
	fmt.Println("This is the difference between using println and printf with Structs object")
	fmt.Printf("Sunny's details are: %+v\n", sunny)
	fmt.Printf("Getting individual value:%v", sunny.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

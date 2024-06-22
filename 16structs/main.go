package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")
	sunny := User{"Sunny", "devsunny@dev.ca", true, 25, 2}
	fmt.Println(sunny)
	fmt.Println("This is the difference between using println and printf with Structs object")
	fmt.Printf("Sunny's details are: %+v \n", sunny)
	fmt.Printf("Getting individual value:%v\n", sunny.Email)
	sunny.GetStatus()
	sunny.NewEmail()
	fmt.Println(sunny)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewEmail() {
	u.Email = "test@hellochanged.ca"
	fmt.Println("Email has been changed :", u.Email)
}

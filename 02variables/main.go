package main

import "fmt"

const LoginToken string = "76576bhgtyvt"

func main() {
	fmt.Println(LoginToken)
	fmt.Printf("Login token is %T \n", LoginToken)

	var username string = "sunny"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isDeveloper bool = true
	fmt.Println(isDeveloper)
	fmt.Printf("Variable is of type: %T \n", isDeveloper)

	var yearsOfExperience uint = 255
	fmt.Println(yearsOfExperience)
	fmt.Printf("Variable if of type: %T \n", yearsOfExperience)

	var height float64 = 5.11
	fmt.Println(height)
	fmt.Printf("Variable if of type: %T \n", height)

	// no var
	numberofProjects := 30000000
	fmt.Println(numberofProjects)

}

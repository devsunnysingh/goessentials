package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Siwtch demo")
	var diceNumber int
	rand.Seed(time.Now().UnixNano())
	diceNumber = rand.Intn(6) + 1
	fmt.Println("Dice number is :", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Hey this is 1")
	case 2:
		fmt.Println("Hey this is 2")
		fallthrough
	case 3:
		fmt.Println("Hey this is 3")
	case 4:
		fmt.Println("Hey this is 4")
		fallthrough
	case 5:
		fmt.Println("Hey this is 5")
		fallthrough
	case 6:
		fmt.Println("Hey this is 6")
	default:
		fmt.Println("Something is wronng")

	}

}

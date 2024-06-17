package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza shop")
	fmt.Println("Please rate our pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating , ", input)
	fmt.Printf("The type of rating is %T", input)
	numNewRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)

	} else {
		numNewRating := numNewRating + 1
		fmt.Println("New rating after addition is:", numNewRating)
	}

}

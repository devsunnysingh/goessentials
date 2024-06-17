package main

import (
	"fmt"
	"sort"
)

// go mod init slices

func main() {
	// we need to initialize if we use this syntax. JHere we initialized it to be empty
	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("type of fruitslist is %T\n", fruitList)
	var fruitList2 = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of Fruitlist2 is %T\n", fruitList2)

	// First we slice a part of the original slice, and then Adding values in slices
	fruitList2 = append(fruitList[1:], "Mango", "Avocado")
	fmt.Println(fruitList2)

	// highscores
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 234
	highScores[3] = 23
	// now whenever we try  doding APPEND, it helps save memory which we might have initialized initially
	highScores = append(highScores, 900)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	// using ... in slices
	var highScoreDemo = []int{123, 334, 445, 555}
	var index int = 2
	fmt.Println(highScoreDemo)
	highScoreDemo = append(highScores[:index], highScoreDemo[index+1:]...)
	fmt.Print(highScoreDemo)

}

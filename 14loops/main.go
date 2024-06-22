package main

import "fmt"

func main() {
	fmt.Println("Loops demo this is")
	days := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	fmt.Println(days)
	for i := 0; i < len(days); i++ {
		fmt.Println("VAlue of ", i+1, ":", days[i])
	}
	// another for type
	// Here day will refer to the index and not the value
	for day := range days {
		fmt.Println(days[day])

	}
	// another for type
	for index, day := range days {
		fmt.Println(index, "is the index and ", day, "is the value")
		fmt.Printf("Index is %v and the value is %v \n", index, day)
		fmt.Println("=============================================")
	}
	// using break and continue
	rogueValue := 1
	for rogueValue < 20 {
		if rogueValue == 15 {
			break
		} else if rogueValue == 10 {
			rogueValue++
			continue
		}
		fmt.Println(rogueValue)
		rogueValue++
	}

}

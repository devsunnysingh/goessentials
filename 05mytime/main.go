package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Uderstanding time function and usage")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("2006-01-02"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println("The date which I created:", createdDate)
}

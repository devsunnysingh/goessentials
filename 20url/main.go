package main

import (
	"fmt"
	"net/url"
)

const urlString string = "https://fakestoreapi.com/products/1"

func main() {
	fmt.Println("Handling URLs")
	fmt.Println(urlString)
	// parsing
	result, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println("Port : ", result.Port())

	qparams := result.Query()
	fmt.Println("The type of query parameters is :", qparams)

}

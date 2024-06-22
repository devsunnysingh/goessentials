package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://fakestoreapi.com/products/1"

func main() {
	fmt.Println("Web Request Demo")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	fmt.Println("Response is of type:", response)
	fmt.Println("\nStatus:", response.Status)
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println("Body of the request:", string(body))
	response.Body.Close()

}

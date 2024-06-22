package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web verb video")
	// PerforGetRequest()
	PerforGetRequestUsingBuilder()
	fmt.Println("-------------------Demo of POST method-------------------")
	PerformPostJSONRequest()
	fmt.Println("-------------------Demo of taking in Form INPUT-------------------")

	PerformPostFormRequest()

}

// Approach 1
func PerforGetRequest() {
	const myUrl = "http://localhost:8000/"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code of the Response : ", response.Status)
	fmt.Println("Content length of the response : ", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

// Approach 2
// Notes: A Builder is used to efficiently build a string  using Write methods.
// Advantage: minimizes memory copying
func PerforGetRequestUsingBuilder() {
	const myUrl = "http://localhost:8000/"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code of the Response : ", response.Status)
	fmt.Println("Content length of the response : ", response.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is  : ", byteCount)
	fmt.Println("Body is : ", responseString.String())
	fmt.Println(string(content))
}

func PerformPostJSONRequest() {
	const myUrl = "http://localhost:8000/post"
	// fake JSON payload
	requestBody := strings.NewReader(`
	{
		"programming_language":"Go",
		"level":"easy",
		"createdBy":"Google"

	}
	`)
	response, errResponses := http.Post(myUrl, "application/json", requestBody)
	if errResponses != nil {
		panic(errResponses)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"
	// formData
	data := url.Values{}
	data.Add("firstname", "sunny")
	data.Add("lastname", "singh")
	data.Add("email", "sunny@loblaws.ca")
	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Data which is sent for POST:", string(content))
}

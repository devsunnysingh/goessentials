package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Testing JSON data with backend")
	EncodeJson()

}

func EncodeJson() {
	sunnyGoCourse := []course{
		{"Go", 20, "Coursera", "abc123", []string{"webdev", "SRE"}},
		{"GCP", 20, "Coursera", "abc123", []string{"DevOPs", "SRE"}},
	}
	finalJSON, err := json.MarshalIndent(sunnyGoCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJSON))

}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Go lang")
	var companyList [10]string
	companyList[0] = "Google"
	companyList[1] = "Microsoft"
	companyList[4] = "CGI"
	fmt.Println(companyList)
	fmt.Println(len(companyList))

	var jobList = [3]string{"consultant", "java dev", "SRE"}
	fmt.Println(jobList)

}

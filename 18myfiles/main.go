package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcomw to doing File Manilpulation")
	content := "This needs to go in the file"
	file, err := os.Create("./myclogfile.txt")
	if err != nil {
		// will show the error and end the program
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is : ", length)
	// Defer is used so that if we want to write any more logic, then we do that and finally defer code will run
	defer file.Close()
	fmt.Println(file.Name())
	readFile(file.Name())

}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}

}

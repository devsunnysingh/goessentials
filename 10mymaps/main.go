// go mod init mymaps
package main
import "fmt"

func main() {
	fmt.Println("Demo of maps concepts in Go lang")
	// languaages:=make()
	languages:=make(map[string]string)
	languages["JS"]="Javascript"
	languages["RB"]="Ruby"
	languages["PY"]="python"

	fmt.Println("List of all languages:",languages)

	fmt.Println("JS stands for :", languages["JS"])
	// deleteing
	delete(languages,"RB");
	fmt.Println("List of all languages:",languages)
// for loop
	for key, value:= range languages {
		fmt.Printf("For key %v, value is %v\n",key,value)
	}

}

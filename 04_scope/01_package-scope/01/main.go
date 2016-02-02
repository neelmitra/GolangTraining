package main

import "fmt"

var x = 42

func main() {
	fmt.Println("Printing from main function : ", x)
	foo()
}

func foo() {
	fmt.Println("Printing from foo function again : ", x)
}

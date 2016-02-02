package main

import "fmt"

//conversion of integer to string , then conversion of string to rune (bytes)

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := "a"
	fmt.Printf("This is %s ", foo)
	fmt.Printf("%T \n", foo)
}

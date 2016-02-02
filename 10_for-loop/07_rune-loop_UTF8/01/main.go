package main

import "fmt"

//conversion of integer to string , then conversion of string to rune (bytes)

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := "a" // string
	boo := 'b' //rune
	doo := `\t` //raw string literal for interpreting characters as it is like \r, \t
	fmt.Printf("This is %s ", foo)
	fmt.Printf("of type %T \n", foo)
	fmt.Printf("This is %v ", boo)
	fmt.Printf("of type %T \n", boo)
	fmt.Printf("This is %v ", doo)
	fmt.Printf("of type %T \n", doo)
}

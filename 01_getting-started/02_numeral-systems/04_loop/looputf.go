package main

import "fmt"

func main() {
	fmt.Println("Printing the Ascii, Binary, Hex and UTF codes from A to Z , in small and caps")
	fmt.Println("Testing Github push")
	for i := 65; i < 127; i++ {
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}
}

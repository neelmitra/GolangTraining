package main

import "fmt"

func main() {
	x := 200

	for i := 5; i < 200; i++ {
		x = x / i
		if x > 1 {
			fmt.Println("The value for x now is", x)
		}
	}

}

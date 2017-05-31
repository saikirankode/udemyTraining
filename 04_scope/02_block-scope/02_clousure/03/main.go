package main

import "fmt"

func main() {
	var x = 0

	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

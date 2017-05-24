package main

import "fmt"

const (
	_ = iota      // 0
	a = iota * 10 //1*10
	b = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}

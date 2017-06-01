package main

import "fmt"

func main() {
	greet("ram", "prasad")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}

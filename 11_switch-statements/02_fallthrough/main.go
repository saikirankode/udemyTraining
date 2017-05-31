package main

import "fmt"

func main() {

	switch "ram prasad" {

	case "sandeep":
		fmt.Println("hello sandeep")
	case "vamsi":
		fmt.Println("hello vamsi")
	case "ram prasad":
		fmt.Println("hello ram prasad")
		fallthrough
	case "tanuj":
		fmt.Println("how are you tanuj")
		fallthrough
	case "arun":
		fmt.Println("whats going on arun")
	default:
		fmt.Println("what have you done this weekend")
	}
}

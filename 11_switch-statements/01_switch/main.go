package main

import "fmt"

func main() {

	switch "kode" {

	case "kiran":
		fmt.Println("Wassup kiran")
	case "kode":
		fmt.Println("Wassup kode")
	case "sai":
		fmt.Println("here we aree sai")
	default:
		fmt.Println("have u have friends")
	}
	// fallthrough (break) is not needed or no default fallthrough
	//fallthrough is optional
	/*
	-- we can specify fallthrough by explicitly stating it
	-- break isn't needed like in other languages
	*/
}

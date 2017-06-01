package main

import "fmt"

func makegreeter() func() string {
	return func() string {
		return "Hello World!"
	}

}

func main() {
	greet := makegreeter()
	fmt.Println(greet())
}

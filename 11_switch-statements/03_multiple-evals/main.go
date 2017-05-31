package main

import "fmt"

func main() {
	switch "ramesh" {
	case "ramesh", "suresh":
		fmt.Println("wassup ramesh, or, err, suresh")
	case "sachin", "kohli":
		fmt.Println("both of them are great batsmen")
	case "abhiruchi", "red saffron":
		fmt.Println("a great collaboration towards success")
	}
}

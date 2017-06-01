package main

import "fmt"

func main() {
	fmt.Println(greet("sandeep", "ramesh"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

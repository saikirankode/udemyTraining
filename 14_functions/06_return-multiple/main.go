package main

import "fmt"

func main() {
	fmt.Println(greet("sai ", "ram "))
}

func greet(fname string, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

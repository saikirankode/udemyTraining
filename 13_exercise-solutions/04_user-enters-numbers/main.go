package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumber int

	fmt.Println("Please enter smallNumber :")
	fmt.Scan(&smallNumber)
	fmt.Println("Please enter largeNumber :")
	fmt.Scan(&largeNumber)

	fmt.Println( largeNumber, "%" ,smallNumber, "=", largeNumber%smallNumber )
}

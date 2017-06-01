package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := " this is accesible within the braces"
		fmt.Println(y)
	}
	//fmt.Println(y)   //outside the scope of y
}

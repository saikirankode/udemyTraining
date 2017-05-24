package main

import "fmt"

func max(x int) int{
	return 42+x
}

func main(){
	max := max(7)
	fmt.Println(max) //max is now result, not a function
}
// this is totally a bad coding this is known as variable shadowing
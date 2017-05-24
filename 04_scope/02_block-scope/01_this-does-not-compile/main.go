package main

import "fmt"

func main(){
	x:=42
	fmt.Println(x)

}

func foo(){

	//no access to x
	//does not compile
	fmt.Println(x)
}

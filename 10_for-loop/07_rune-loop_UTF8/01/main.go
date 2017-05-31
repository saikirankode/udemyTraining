package main

import "fmt"

func main() {
	for i := 50; i <= 150; i++ {
		fmt.Println('i', " - ", string(i), " - ", []byte(string(i)))
	}
	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))

	fool := "b"
	fmt.Println(fool)
	fmt.Printf("%T \n", fool)

}

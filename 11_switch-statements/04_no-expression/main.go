package main

import "fmt"

func main() {

	myFriendName := "sai"

	switch {
	case len(myFriendName) == 3:
		fmt.Println("hai, hello friend with name length 3")
	case myFriendName == "sai":
		fmt.Println("dont be affraid")
	case myFriendName == "sunny":
		fmt.Println("get well soon sunny")
	case myFriendName == "sateesh", myFriendName == "satya":
		fmt.Println("both of them are bavas for me")
	case myFriendName == "sana":
		fmt.Println("my cute little sana")
	case myFriendName == "shrey":
		fmt.Println("when are you coming nj shrey")
	case myFriendName == "sai":
		fmt.Println("work hard sai")
	default:
		fmt.Println("u have less time, dont waste")
	}
}

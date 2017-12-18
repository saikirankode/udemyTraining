package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says,"good morning, james."`)
}
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says,"shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human){
	h.speak()
}

func main() {

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)

	m := map[string]int{
		"kode":  20,
		"sai":   23,
		"kiran": 24,
	}
	fmt.Println(m)

	p1 := person{
		"saikiran",
		"kode",
	}
	//p1.speak()

	for _, x := range xi {
		fmt.Println(x)
	}

	sa1 := secretAgent{
		person{
			"intallent",
			"inc",
		},
		true,
	}
	//sa1.speak()
	//sa1.person.speak()
	saySomething(p1)
	saySomething(sa1)
}

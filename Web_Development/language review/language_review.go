package language_review

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func(p person)speak(){
	fmt.Println(p.fname, p.lname, `says, "Good morning, Bava."`)
}

func(sa secretAgent)speak(){
	fmt.Println(sa.fname, sa.lname, `says, "Shakken, but not strirred."`)
}

type human interface{
	speak()
}

func saySomething(h human){
	h.speak()
}
func main() {
	/*sl := []int{4, 58, 6, 3, 6}
	fmt.Println(sl)

	m:= map[string]int{
		"Todd":45,
		"satya":32,
		}
	fmt.Println(m)*/

	p1:= person{
		"saikiran",
		"kode",
		}
	//p1.speak()

	sa1:= secretAgent{
			person{
				"james",
				"bond",
			},
			true,
	}
	/*sa1.speak()
	sa1.person.speak()*/
	saySomething(p1)
	saySomething(sa1)

}

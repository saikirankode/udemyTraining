package main

import (
	"html/template"
	//"os"
	"log"

	//"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("contact.html"))
}

type address struct {
	Addressline1 []string
	Addressline2 []string
}

func main() {
	add := address{
		Addressline1: []string{"Ward Edwards 1400,", "Warrensburg, MO 64093", "phone: 1-877-729-8266", "660-543-4111", "email: admit@ucmo.edu"},
		Addressline2: []string{"Vignana Jyothi Group", "of Institutions,", "H.No.7-1-4, Begumpet", "Hyderabad-500 016.", "Phone : 040-23731555"},
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, add)

	if err != nil {
		log.Fatalln(err)
	}
}

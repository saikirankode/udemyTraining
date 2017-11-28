package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}

	err = tpl.Execute(file, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

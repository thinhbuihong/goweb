package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.go.html")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error create file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, "thinh")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, "thinh")
	if err != nil {
		log.Fatalln(err)
	}
}

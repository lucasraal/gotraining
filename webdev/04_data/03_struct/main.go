package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name	string
	Motto	string
}

func init()  {
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.ParseFiles("tpl2.gohtml"))
}

func main()  {
	sages := sage{
		Name: "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
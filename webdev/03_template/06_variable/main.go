package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("one.gohtml"))
}

func main()  {
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", 
	"you gotta believe!")
	if err != nil {
		log.Fatalln(err)
	}
}
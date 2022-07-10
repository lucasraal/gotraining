package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main()  {
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", 
	"this text will replace the {{.}}")
	if err != nil {
		log.Fatalln(err)
	}
}
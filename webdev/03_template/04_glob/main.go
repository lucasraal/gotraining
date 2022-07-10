package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main()  {
	
	// parsing all files from a folder as templates
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// serving a template
	// (it just executes the first template "one.gohtml")
	fmt.Println("Running Template")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// serving selected template
	fmt.Println("Running Template - two.gohtml")
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
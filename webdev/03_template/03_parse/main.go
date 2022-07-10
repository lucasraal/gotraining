package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func main()  {
	//creating text files
	author := `Copyright Lucas`

	one := `
	***
	ONE ` + author + `
	***
	`
	nf_one, err := os.Create("one.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf_one.Close()
	io.Copy(nf_one, strings.NewReader(one))

	two := `
	***
	TWO ` + author + `
	***
	`
	nf_two, err := os.Create("two.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf_two.Close()
	io.Copy(nf_two, strings.NewReader(two))

	// creating template
	tpl, err := template.ParseFiles("one.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// serving a template
	fmt.Println("Running Template - First Time")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// parsing one more template
	tpl, err = tpl.ParseFiles("two.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// serving a template 
	// (it just executes the first defined template "one.txt")
	fmt.Println("Running Template - Second Time")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// serving selected template
	fmt.Println("Running Template - one.txt")
	err = tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Running Template - two.txt")
	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
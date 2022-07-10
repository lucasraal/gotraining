package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func dayMonthYear(t time.Time) string {
	// 02 is day; 01 is month; 06 is year; 3,4,5 is HH,MM,SS; 7 is GMT
	return t.Format("02/01/2006")
}

func kitchenFmt(t time.Time) string {
	// using constants as predefined formats
	return t.Format(time.Kitchen)
}

func fullDate(t time.Time) string {
	// using constants as predefined formats
	return t.Format(time.RFC850)
}

func personalDate(t time.Time) string {
	// 02 is day; 01 is month; 06 is year; 3,4,5 is HH,MM,SS; -7 is GMT
	return t.Format(
		`Today is day 02, month 01, and year 2006.
The time is 03:04 GMT -07`,
	)
}

var fm = template.FuncMap{
	"fdateMDY":		dayMonthYear,
	"kitchen":		kitchenFmt,
	"fulldateFmt":	fullDate,
	"personalFmt":	personalDate,
}

func init()  {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main()  {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
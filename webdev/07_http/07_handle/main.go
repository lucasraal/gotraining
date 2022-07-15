package main

import (
	"io"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Using standard mux from http package\nDogs are our best friends")
}

type cat int

func (c cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Using standard mux from http package\nCats have their own personality")
}

func main()  {
	var d dog
	var c cat

	http.Handle("/dog/", d)
	http.Handle("/cat/", c)

	// using nil will give me the standard mux from http package
	http.ListenAndServe(":8080", nil)
}
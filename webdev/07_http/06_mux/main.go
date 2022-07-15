package main

import (
	"io"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dogs are our best friends")
}

type cat int

func (c cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Cats have their own personality")
}

func main()  {
	var d dog
	var c cat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)

	http.ListenAndServe(":8080", mux)
}
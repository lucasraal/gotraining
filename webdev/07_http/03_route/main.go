package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the landing page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dogs are our best friend")
}

func cat(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Cats have their own personality")
}

func main()  {
	port := ":8080"
	fmt.Fprintf(os.Stdout, "Starting server on port %v\n", port)
	http.HandleFunc("/", home)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)
	http.ListenAndServe(port, nil)
}
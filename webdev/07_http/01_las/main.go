package main

import (
	"fmt"
	"net/http"
)

// func ListenAndServe(addr string, handler Handler) error

// A type Handler is an interface to every type that has
// the method ServeHTTP(ResponseWriter, *Request)
// ServeHTTP is the identifier and could be any name
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type helloHandler string

func (h helloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintln(w, "Hello, World!")
}

func main()  {
	var hello helloHandler
	port := ":8081"
	fmt.Printf("Starting server on localhost%v", port)
	http.ListenAndServe(port, hello)
}
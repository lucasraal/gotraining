package main

import (
	"fmt"
	"net/http"
)

type handler string

func (h handler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	// you can find req values on
	// https://pkg.go.dev/net/http#Request
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(res, 
			`Method: %v
			Headers: %v
			Body: %v`,
			req.Method,
			req.Header,
			req.Body,
		)
	case "/dog/":
		fmt.Fprintln(res, "Dogs are our best friends")
	case "/cat/":
		fmt.Fprintln(res, "Cats have their own personality")
	}
}

func main()  {
	var h handler
	http.ListenAndServe(":8080", h)
}
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
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
	port := ":8080"
	
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {		
		http.ListenAndServe(port, hello)
		// code never gets here below as it does not stop listening
		defer wg.Done()
		time.Sleep(time.Second * 10)
		fmt.Println("\nSERVER CLOSED")
	}()

	fmt.Printf("Server running on port %v", port)
	wg.Wait()
}
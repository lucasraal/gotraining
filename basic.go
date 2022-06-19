package main

import (
	"fmt"
)

func basic() {

	s := "Hello, World!"
	cs := []byte(s)

	fmt.Println(cs)
	fmt.Printf("%T\n", cs)
}
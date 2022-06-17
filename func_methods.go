package main

import (
	"fmt"
)

type person struct {
	first	string
	last	string
}

type secretAgent struct {
	person
	ltk		bool
}

// func (r receiver) identifier(parameters) (return) {code}
// if I define a receiver, it means that the function gets attached to any value of the receiver type

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
	if s.ltk {
		fmt.Println("I have license to kill")
	} else {
		fmt.Println("I do not have license to kill")
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Lucas",
			last: "Lima",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Lunna",
			last: "Haullier",
		},
		ltk: false,
	}

	fmt.Println(sa1)
	// as func speak is defined as a method to secretAgent, i don't have to write a function and use sa1 as an argument
	// instead, I can just call the function speak() as a method (like an object) of sa1
	sa1.speak()

	fmt.Println(sa2)
	sa2.speak()
}
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

// funf (r receiver) identifier(parameters) (return) {code}
// if I define a receiver, it means that the function gets attached to any value of the receiver type

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
	if s.ltk {
		fmt.Println("I have license to kill", "- said the secretAgent")
	} else {
		fmt.Println("I do not have license to kill", "- said the secretAgent")
	}
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
	fmt.Println("I do not know whether I have license to kill", "- said the person")
}

type human interface {
	speak()
}

// If I want to use specific "properties or objetcs" from that struct/type converted using interface
// then I have to use a switch statement on type and describe which "objetcs" I would like to use for each type
func foo(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am the person", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("I am the secretAgent", h.(secretAgent).first, h.(secretAgent).last)
	}
}

func bar(h human) {
	fmt.Println("I am a person", h)
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

	p1 := person{
		first: "Alberto",
		last: "Penzin",
	}

	fmt.Println(sa1)
	// as func speak is defined as a method to secretAgent, i don't have to write a function and use sa1 as a parameter
	// instead, I can just call the function speak() as a method (like an object) of sa1
	sa1.speak()

	fmt.Println(sa2)
	sa2.speak()

	fmt.Println(p1)
	p1.speak()

	foo(sa1)
	foo(sa2)
	foo(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
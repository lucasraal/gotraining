package main

import "fmt"

type person struct {
	first	string
	last	string
	age		int
}

func main() {
	p1 := person{"Lucas", "Lima", 29}
	fmt.Println(p1)
	p2 := person{
		first: "Antonio",
		last: "Lima",
		age: 30,
	}
	fmt.Println(p2.first)
}
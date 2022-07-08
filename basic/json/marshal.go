package main

import (
	"encoding/json"
	"fmt"
)

// uppercase identifier if you want it to be exported
type person struct {
	First	string
	Last	string
	Age		int
}

func main()  {
	p1 := person{
		First: "James",
		Last: "Bond",
		Age: 32,
	}

	p2 := person{
		First: "Mona",
		Last: "Lisa",
		Age: 22,
	}

	people := []person{p1, p2}

	fmt.Println(people)
	fmt.Printf("%T\n", people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(string(bs))
}
package main

import (
	"encoding/json"
	"fmt"
)

// the tag allows to define the field name in the JSON string
// whereas you can use a different field name in Go struct
// omitempty is an option that allows you to ommit empty values in the JSON string
type person struct {
	NameF string `json:"First,omitempty"`
	NameL  string `json:"Last,omitempty"`
	Age   int    `json:"Age,omitempty"`
	Salary	int	
}

func main()  {

	s := `[{"First":"Zhong","Last":"Li","Age":32},{"First":"Ast","Last":"Mona","Age":21},{"First":"Ningguang","Last":"Liyue","Age":24}]`

	// convert the JSON to byte
	bs := []byte(s)

	// have to create a slice of the same struct as is the JSON
	// people := []person{}
	// the var is more readable - write like a master
	var people []person

	// gotta use pointer to change the value of the slice from JSON to Go struct
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	// now I am able to manipulate the slice in my code
	fmt.Println(people[0].NameF)
	fmt.Printf("%T\n",people[0].NameF)
	people[0].NameF = "Lucas"
	fmt.Println(people[0].NameF)
	fmt.Println(people)

	for i := range people {
		people[i].Salary = 3200
	}
	fmt.Println(people)
}
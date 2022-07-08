package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name	string
	Age		int
	Balance	float64
}

func main() {
	p1 := person{"Lucas", 29, 9560.20}
	p2 := person{"Lunna", 22, 18540.80}
	p3 := person{"Mariana", 25, 850.30}
	p4 := person{"Alana", 20, 42.60}

	people := []person{p1, p2, p3, p4}

	// fmt.Println(people)

	// sort.SliceStable preserves the original order for equal elements
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	// fmt.Println("By age:", people)

	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	// fmt.Println("By name, age:", people)

	// by balance ">" means higher first (descending)
	sort.SliceStable(people, func(i, j int) bool { return people[i].Balance > people[j].Balance })
	// fmt.Println("By balance, name, age:", people)

	fmt.Println("\nRICHEST PEOPLE RANKING")
	for i, v := range people {
		fmt.Println("\n--- RANKING NUMBER", i+1, "---")
		fmt.Println("Name:", v.Name)
		fmt.Println("Age:", v.Age)
		fmt.Println("Balance:", v.Balance)
	}
	fmt.Println("\n----------------------")
}
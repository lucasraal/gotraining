package main

import (
	"fmt"
)

var name string = "Lucas"
var age int = 29
var isStudent bool = true
var bank_balance float64 = 2365.22

func main() {

	status := ""
	travel := "CANNOT"

	if isStudent {
		status = "is currently enrolled in TUM"
	} else {
		status = "is not studying yet"
	}

	if bank_balance > 2000 {
		travel = "CAN"
	}

	n, err := fmt.Println("Name:", name, "Age:", age)
	fmt.Println(n)
	fmt.Println(err)

	fmt.Println(name, status)

	fmt.Printf("%v, your bank account balance is currently: $%v\nYou %v travel\n", name, bank_balance, travel)
}


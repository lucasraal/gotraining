package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println(x)
	fmt.Println(&x)
	foo(&x)
	fmt.Println(x)
	fmt.Println(&x)
}

// by using pointers, I don't have to return a new value from a function and
// and assign this value to the original variable
// I can just locate where the variable's value is stored in the memory and
// and change the value stored in that part of the memory
func foo(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
	*x = 22
	fmt.Println(*x)
}
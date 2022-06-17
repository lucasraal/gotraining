package main

import (
	"fmt"
)

func main() {
	// use () after writing the anonymous function in order to run it
	func () {
		fmt.Println("This is an anonymous function")
	}()

	// you can also use arguments when calling an anonymous function
	func (x int)  {
		fmt.Println("Anonymous function with a parameter, which current argument is", x)
	}(33)

	// you can assign a function to a variable as well
	f := func () {
		fmt.Println("This function has been assigned to a variable")
	}
	f()

	g := func (x int) {
		fmt.Println("The argument value is", x)
	}
	g(24)

	// now we call the function that returns a function and we run the function that has been returned
	fmt.Println(foo()())
}

// let's build a function that returns a function that returns an integer value
func foo() (func() int) {
	return (func() int {
		return 451
	})
}
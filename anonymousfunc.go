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
}
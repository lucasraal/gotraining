package main

import "fmt"

func main() {
	x := []int{1, 5, 3}
	y := sum(x...)
	fmt.Println("the total is", y)
}

func sum(x ...int) (int) {

	y := 0
	fmt.Println(y)

	for _, j := range x {
		y += j
		fmt.Println(y)
	}

	return y
}
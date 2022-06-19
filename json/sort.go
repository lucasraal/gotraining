package main

import (
	"fmt"
	"sort"
)

func main()  {
	xi := []int{5, 9, 8, 6, 2, 4, 3, 7, 1, 0}
	xs := []string{"Lucas", "Isabela", "Lunna", "Alana", "Mariana", "Sabrina"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
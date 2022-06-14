package main

import "fmt"

func main() {
	x := []int{2, 5, 3+4, 9}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	// The for statement supports one additional form that uses the keyword range
	for i, v := range x {
		fmt.Println(i, v)
	}
	// the traditional way of looping through a slice is
	for ind := 0; ind < len(x); ind++ {
		fmt.Println(ind, x[ind], "oi")
	}

	x = append(x, 10, 11)
	fmt.Println(x)

	y := []int{21, 22, 23, 24}
	x = append(x, y...) // the ... is a variadict parameter that "returns" all elements from y
	fmt.Println(x)
	// x = append(x, y) doesn't work since x elements are int, but y is a slice, not an integer type element

	// deleting elements from a slice
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	a = append(a[:3], a[5:]...) // deleted element on position 3 and 4
	fmt.Println(a)

	// when you make a slice with a certain capacity
	// the runtime doesn't have to create a new array every time you append a new element
	aa := make([]int, 2, 4)
	fmt.Println(aa, len(aa), cap(aa))
	// if appending a new element exceeds the cap(), the runtime will double the cap()
	// in this case the runtime 
	aa = append(aa, 1, 1, 2)
	fmt.Println(aa, len(aa), cap(aa))

	nn := []string{"Lucas", "Lima", "29", "Brazil"}
	mm := []string{"Lunna", "Haullier", "22", "England"}
	fmt.Println(nn, mm)
	xp := [][]string{nn, mm}
	fmt.Println(xp)
	fmt.Println(xp[0][3])

	// MAP
	hhh := map[string]int{
		"Lucas":	29,
		"Mariana":	20,
	}
	fmt.Println(hhh)
	// use make() to create an empty map
	sss := make(map[string]int)
	sss["Isabela"] = 25
	sss["Aurora"] = 27
	fmt.Println(sss)

	// comma ok expression to check whether 0 is a false value or 0 is a true value
	vvv := sss["Aurora"]
	if _, ok := sss["Aurora"]; ok {
		fmt.Println(vvv)		
	}

	// for-range loop on a map
	for iii, zzz := range sss {
		fmt.Println(iii, zzz)		
	}
}
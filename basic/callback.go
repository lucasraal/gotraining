package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println("The sum of ALL numbers is", sum(x...))
	fmt.Println("The sum of EVEN numbers is", evenSum(sum, x...))
	fmt.Println("The sum of ODD numbers is", oddSum(sum, x...))
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func evenSum(f func(x ...int) int, y ...int) int {
	even := []int{}
	for _, v := range y {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
//	total := f(even...)
//	return total
	return f(even...)
}

func oddSum(f func(x ...int) int, y ...int) int {
	odd := []int{}
	for _, v := range y {
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
//	total := f(odd...)
//	return total
	return f(odd...)
}
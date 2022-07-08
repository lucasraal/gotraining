package main

import "fmt"

func main()  {
	odd := make(chan int)
	even := make(chan int)
	finish := make(chan bool)

	go f_odd_even(odd, even, finish)

	func () { 
		for {
			select {
			case v := <-odd:
				fmt.Println("Odd value:", v)
			case v := <-even:
				fmt.Println("Even value:", v)
			case <-finish:
				fmt.Println("--- INFINITE LOOP FINISHED ---")
				return
			}
		}
	}()

	fmt.Println("--- PROCESS FINISHED ---")
}

func f_odd_even(o, e chan<- int, f chan<- bool)  {
	for i := 1; i <= 10; i++ {
		if i % 2 != 0 {
			o<- i
		} else {
			e<- i
		}
	}
	// close(o)
	// close(e)
	f<- true
}
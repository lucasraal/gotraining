package main

import "fmt"

func main()  {
	runNoBuffer()
	runBuffer()
	runDoubleBuffer()
}

func runNoBuffer()  {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func runBuffer()  {
	c := make(chan int, 1)
	c<- 42
	fmt.Println(<-c)
}

func runDoubleBuffer()  {
	c := make(chan int, 2)
	c<- 25
	c<- 50
	fmt.Println(<-c)
	fmt.Println(<-c)
}
package main

import "fmt"

func main()  {
	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("Channel transfered data from a concurrent routine to another")
}

func send(cs chan<- int)  {
	cs<- 42
	fmt.Println("Sent", cs)
}

func receive(cr <-chan int)  {
	fmt.Println("Received", cr)
	fmt.Println("Received", <-cr)
}
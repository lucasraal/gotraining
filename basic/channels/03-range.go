package main

import "fmt"

func main()  {
	c := make(chan int)	

	go sender(c)

	receiver(c)

	fmt.Println("--- PROCESS FINISHED ---")
}

func sender(cs chan<- int)  {
	for i := 0; i < 10; i++ {
		cs<- i
	}
	// if I don't close the channel, the range loop below will keep waiting data from this channel
	// while it has already stopped having values assigned to it
	// thus giving rise to a deadlock condition
	close(cs)
}

func receiver(cr <-chan int)  {
	for v := range cr {
		fmt.Println("Value received:", v)
	}
}
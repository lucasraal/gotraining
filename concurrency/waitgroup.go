package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// FOO WILL RUN IN PARALEL WITH THE CURRENT CODE, as we added "go" before "foo()", thus it will run as a different routine
	// However, we add it to the wait group because we need to instruct 
	// the runtime to wait for the complete execution of foo (running in an alternative routine)
	// before terminating the func main() (which is running in the main rountine)
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	
	fmt.Println("Waiting foo to terminate the function...")
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("--- FUNCTION TERMINATED ---")
}

func foo() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("foo:", i)
	}
	time.Sleep(1000 * time.Millisecond)
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("bar:", i)
	}
}
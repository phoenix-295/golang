package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("Arch", runtime.GOARCH)
	fmt.Println("CPU", runtime.NumCPU())
	wg.Add(1)
	go foo()
	bar()
	wg.Wait()
	fmt.Println("Routines", runtime.NumGoroutine())
}

func foo() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Bar: ", i)
	}
}

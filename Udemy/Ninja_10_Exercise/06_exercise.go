package main

import (
	"fmt"
)
func main() {
	num := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			num <- i
		}
		close(num)
	}()
	rev(num)
}

func rev(x <-chan int) {
	for i := range x {
		fmt.Println(i)
	}
}

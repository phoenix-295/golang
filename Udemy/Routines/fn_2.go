package main

import "fmt"

func main() {
	e := make(chan int)
	o := make(chan int)
	go eve(e, o)
	rec(e, o)
}

func eve(e, o chan<- int) {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func rec(e, o <-chan int) {
	for r := range e {
		fmt.Println(r)
	}
	for i := range o {
		fmt.Println(i)
	}
}

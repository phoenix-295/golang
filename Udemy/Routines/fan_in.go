package main

import "fmt"

func main() {
	ev := make(chan int)
	// od = make(chan int)

	go getev(ev)
	rec(ev)

}

func getev(ev chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			ev <- i
		}
	}
	close(ev)
}

func rec(r <-chan int) {
	for rn := range r {
		fmt.Println(rn)
	}
}

package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}

	}()

	go func() {
		for i := 11; i <= 20; i++ {
			ch <- i
		}

	}()
	go func() {
		for i := 21; i <= 30; i++ {
			ch <- i
		}

	}()
	go func() {
		for i := 31; i <= 40; i++ {
			ch <- i
		}

	}()
	go func() {
		for i := 41; i <= 50; i++ {
			ch <- i
		}

	}()
	go func() {
		for i := 51; i <= 60; i++ {
			ch <- i
		}

	}()
	go func() {
		for i := 61; i <= 70; i++ {
			ch <- i
		}

	}()
	go func() {
		for i := 71; i <= 80; i++ {
			ch <- i
		}

	}()
	go func() {
		for i := 81; i <= 90; i++ {
			ch <- i
		}

	}()
	go func() {
		for i := 91; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()
	rec(ch)
}

func rec(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

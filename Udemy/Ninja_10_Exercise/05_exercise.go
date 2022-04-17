package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(){
		for i:=1;i<=5;i++{
			c <- i
		}
	}()

	v, ok := <- c
	fmt.Println(v, ok)

	close(c)
	
	v, ok = <- c
	fmt.Println(v, ok)
}

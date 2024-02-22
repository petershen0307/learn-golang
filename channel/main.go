package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go work(ch)
	go work(ch)
	ch <- 1
	ch <- 2
	close(ch)
	fmt.Println("channel closed")
	time.Sleep(2 * time.Second)

	/*
		the output will be like this
			channel closed
			1 true
			2 true
			0 false
			0 false
	*/
}

func work(ch chan int) {
	time.Sleep(1 * time.Second)
	v, ok := <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
}

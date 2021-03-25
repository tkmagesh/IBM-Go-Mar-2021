package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quitting")
			return
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan int)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(<-c)
		}
	}()
	go fibonacci(c, q)
	fmt.Println("Hit Enter if you want to quit")
	var input string
	fmt.Scanln(&input)
	q <- 0
	//time.Sleep(5 * time.Second)
}

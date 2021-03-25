package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	go func() {
		fmt.Println("Writing 10")
		ch <- 10
		fmt.Println("Writing 20")
		ch <- 20
		fmt.Println("Writing 30")
		ch <- 30
		fmt.Println("Writing 40")
		ch <- 40
		fmt.Println("Writing 50")
		ch <- 50
		fmt.Println("Writing 60")
		ch <- 60
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

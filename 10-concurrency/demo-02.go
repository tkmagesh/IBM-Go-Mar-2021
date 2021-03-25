package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("Writing data into the channel")
		//ch <- 10
	}()
	fmt.Println("Reading data from the channel")
	data := <-ch
	fmt.Println(data)
}

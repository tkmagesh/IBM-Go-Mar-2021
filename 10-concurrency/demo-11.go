package main

import (
	"fmt"
	"time"
)

/* func fn(ch1 chan int, ch2 chan int) {
	n1 := <-ch1
	fmt.Printf("n1 = %d\n", n1)
	n2 := <-ch2
	fmt.Printf("n2 = %d\n", n2)
} */

func fn(ch1 chan int, ch2 chan int) {
	for {
		select {
		case n1 := <-ch1:
			fmt.Printf("n1 = %d\n", n1)
		case n2 := <-ch2:
			fmt.Printf("n2 = %d\n", n2)
		}
	}
}

func f1(ch1 chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Writing data into ch1")
	ch1 <- 10
}

func f2(ch2 chan int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Writing data into ch2")
	ch2 <- 20
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go fn(ch1, ch2)
	go f1(ch1)
	go f2(ch2)
	fmt.Println("Press ENTER to continue...")
	var input string
	fmt.Scanln(&input)
}

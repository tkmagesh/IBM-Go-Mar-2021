package main

import (
	"fmt"
)

func print(s string, ch chan string) {
	fmt.Println(s)
	ch <- s + " - done"
}

func main() {
	ch := make(chan string)
	go print("World", ch)
	go print("Hello", ch)
	go print("Magesh", ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("end of main")
}

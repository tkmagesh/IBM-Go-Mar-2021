package main

import (
	"fmt"
	"time"
)

func print(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {

	go print("World")
	go print("Hello")
	//go print("Magesh")
	var input string
	fmt.Scanln(&input)
	fmt.Println("end of main")
}

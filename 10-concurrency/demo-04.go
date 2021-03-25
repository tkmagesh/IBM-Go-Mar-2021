package main

import (
	"fmt"
	"time"
)

func print(s string) {
	fmt.Println(s)

}

func main() {

	go print("World")
	go print("Hello")
	go print("Magesh")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("end of main")
}

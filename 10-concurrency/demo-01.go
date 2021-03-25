package main

import (
	"fmt"
)

func print(s string) {
	fmt.Println(s)
}

func main() {
	go print("Hello")
	print("World")
	fmt.Println("end of main")
}

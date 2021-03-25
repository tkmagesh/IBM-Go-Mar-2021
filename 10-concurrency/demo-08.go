package main

import (
	"fmt"
	"sync"
	"time"
)

var data string
var mutex = &sync.Mutex{}

func print(str string) {
	for {
		if data != str {
			mutex.Lock()
			fmt.Println(str)
			data = str
			mutex.Unlock()
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	go print("Hello")
	go print("World")
	var input string
	fmt.Scanln(&input)
}

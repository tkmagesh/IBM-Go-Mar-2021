package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	stop := time.After(5 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-stop:
			fmt.Println("Done")
			return
		default:
			time.Sleep(50 * time.Millisecond)
			fmt.Print(".")
		}
	}
}

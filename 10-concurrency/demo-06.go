package main

import (
	"fmt"
	"time"
)

func print(s string, c1 chan string, c2 chan string) {
	for i := 0; i < 5; i++ {
		<-c1
		fmt.Println(s)
		time.Sleep(1000 * time.Millisecond)
		c2 <- "done"
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go print("World", c1, c2)
	go print("Hello", c2, c1)
	c2 <- "start"
	var input string
	fmt.Scanln(&input)
	fmt.Println("end of main")
}

/* Modify the above code in such a way that it prints the following */
/*
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/

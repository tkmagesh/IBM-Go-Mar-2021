package main

import "fmt"

func main() {
	f1()
}

func f1() {
	fmt.Println("f1 initiated")
	defer f2()
	fmt.Println("f1 exits")
	return

}

func f2() {
	fmt.Println("f2 triggered")
}

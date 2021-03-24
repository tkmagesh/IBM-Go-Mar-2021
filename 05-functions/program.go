package main

import (
	"errors"
	"fmt"
)

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(100, 200))

	/* divide := func(x, y int) (int, int) {
		quotient := x / y
		remainder := x % y
		return quotient, remainder
	} */

	divide := func(x, y int) (quotient, remainder int, err error) {
		if y == 0 {
			err = errors.New("Invalid arguments")
			return
		}
		quotient = x / y
		remainder = x % y
		return
	}

	/* quotient, remainder, err := divide(100, 0)
	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err)
	} else {
		fmt.Println(quotient, remainder)
	} */

	_, _, err := divide(100, 7)
	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err)
	} else {
		fmt.Println("Operation successful")
	}

	//anonymous function
	(func() {
		fmt.Println("Anonymous function executed")
	})()

	up, down := getCounter()
	fmt.Println(up()) //=> 1
	fmt.Println(up()) //=> 2
	fmt.Println(up()) //=> 3
	fmt.Println(up()) //=> 4

	fmt.Println(down()) //=> 3
	fmt.Println(down()) //=> 2
	fmt.Println(down()) //=> 1
	fmt.Println(down()) //=> 0
	fmt.Println(down()) //=> -1

	fmt.Println(aggregate(10, 20))
	fmt.Println(aggregate(10, 20, 30, 40, 50))

	nos := []int{3, 5, 7, 9}
	fmt.Println(aggregate(nos...))
}

func aggregate(nos ...int) int {
	result := 0
	for _, n := range nos {
		result += n
	}
	return result
}

//closures

func getCounter() (func() int, func() int) {
	var count int = 0

	up := func() int {
		count += 1
		return count
	}
	down := func() int {
		count -= 1
		return count
	}
	return up, down
}

package main

import (
	"errors"
	"fmt"
	"strconv"
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
	fmt.Println(aggregate(10))
	fmt.Println(aggregate())
	fmt.Println(aggregate(10, 20, 30, 40))
	fmt.Println(aggregate(10, 20, "30", 40))
	fmt.Println(aggregate(10, 20, "30", 40, "abc"))
	fmt.Println(aggregate(10, 20, []int{30, 40}))
	fmt.Println(aggregate(10, "20", []interface{}{30, "40"}))
}

/*
func aggregate(nos ...int) int {
	result := 0
	for _, n := range nos {
		result += n
	}
	return result
}
*/

/* func aggregate(nos ...interface{}) int {
	result := 0
	for _, no := range nos {
		value, ok := no.(int)
		if ok {
			result += value
		}
		value2, ok := no.(string)
		if ok {
			if value, err := strconv.Atoi(value2); err == nil {
				result += value
			}
		}
		if list, ok := no.([]int); ok {
			intfList := make([]interface{}, len(list))
			for idx, v := range list {
				intfList[idx] = v
			}
			result += aggregate(intfList...)
		}
		if list, ok := no.([]interface{}); ok {
			result += aggregate(list...)
		}
	}
	return result
} */

func aggregate(nos ...interface{}) int {
	result := 0
	for _, no := range nos {
		switch no.(type) {
		case int:
			result += no.(int)
		case string:
			if val, err := strconv.Atoi(no.(string)); err == nil {
				result += val
			}
		case []int:
			numList := no.([]int)
			intfList := make([]interface{}, len(numList))
			for idx, v := range numList {
				intfList[idx] = v
			}
			result += aggregate(intfList...)

		case []interface{}:
			result += aggregate(no.([]interface{})...)
		}
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

getCounter()

func fn(){
	i := 10
}

fn()

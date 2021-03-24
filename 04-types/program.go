package main

import "fmt"

func main() {

	//type conversion
	var i int = 50
	var f float64
	f = float64(i)
	fmt.Printf("%v\n", f)

	//type assertion
	var x interface{}
	x = 10
	//x = "Magesh"
	fmt.Printf("%v\n", x)
	value, ok := x.(string)
	if ok {
		fmt.Printf("%v is type string\n", value)
	}

	var nos interface{}
	nos = []int{1, 2, 3, 4, 5}
	_, ok = nos.([]int)
	if ok {
		fmt.Printf("%v is int slice\n", nos)
	}
	fmt.Printf("typeof f = %T, x = %T, nos = %T\n", f, x, nos)
}

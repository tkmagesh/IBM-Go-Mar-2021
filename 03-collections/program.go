package main

import "fmt"

/*
Array - fixed size
Slice
Range
Map (key,value)
*/
func main() {
	//Array
	var nos [10]int
	for i := 0; i < 10; i++ {
		nos[i] = i * 2
	}
	for i := 0; i < 10; i++ {
		fmt.Println(nos[i])
	}

	/* strs := [2]string{"Hello", "World"} */
	strs := [...]string{"Hello", "World"}
	//fmt.Println(strs)
	fmt.Printf("%q\n", strs)

	//multi dimensional array
	var matrix [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = fmt.Sprintf("row - %d coloum - %d", i, j)
		}
	}
	fmt.Printf("%q\n", matrix)

}

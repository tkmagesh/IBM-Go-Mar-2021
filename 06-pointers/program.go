package main

import "fmt"

func main() {
	/* no := 10
	addressOfNo := &no */
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr, *noPtr)
	n1, n2 := 10, 20
	fmt.Printf("Before swapping n1=%d n2=%d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping n1=%d n2=%d\n", n1, n2)

	nos := []int{10, 20, 30}
	fmt.Printf("Initial list %v\n", nos)
	insert(&nos, 100)
	fmt.Printf("After appending 100, list = %v\n", nos)
}

func insert(listPtr *[]int, no int) {
	*listPtr = append(*listPtr, no)
}

func swap(n1, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
	fmt.Printf("[@swap] after swapping n1=%d n2=%d\n", *n1, *n2)
}

func swap2(n1, n2 int) (int, int) {
	return n2, n1
}

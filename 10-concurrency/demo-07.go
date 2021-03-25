package main

import "fmt"

/*
func sum(nos []int) int {
	result := 0
	for _, n := range nos {
		result += n
	}
	return result
}
func main() {
	nos := []int{4, 1, 5, 2, 6, 3, 7, 9, 8, 10, 14, 16}
	firstSet := nos[:len(nos)/2]
	secondSet := nos[len(nos)/2:]
	firstResult := sum(firstSet)
	secondResult := sum(secondSet)
	result := firstResult + secondResult
	fmt.Printf("result = %d\n", result)
}
*/

func sum(nos []int, resultChan chan int) {
	result := 0
	for _, n := range nos {
		result += n
	}
	resultChan <- result
}

//using 2 channels
/* func main() {
	nos := []int{4, 1, 5, 2, 6, 3, 7, 9, 8, 10, 14, 16}
	firstSet := nos[:len(nos)/2]
	secondSet := nos[len(nos)/2:]
	firstChan := make(chan int)
	secondChan := make(chan int)
	go sum(firstSet, firstChan)
	go sum(secondSet, secondChan)
	result := <-firstChan + <-secondChan
	fmt.Printf("result = %d\n", result)
} */

func main() {
	nos := []int{4, 1, 5, 2, 6, 3, 7, 9, 8, 10, 14, 16}
	firstSet := nos[:len(nos)/2]
	secondSet := nos[len(nos)/2:]
	resultChan := make(chan int)
	go sum(firstSet, resultChan)
	go sum(secondSet, resultChan)
	result := <-resultChan + <-resultChan
	fmt.Printf("result = %d\n", result)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func goRoutine1(wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Routine 1 completed")
	wg.Done()
}

func goRoutine2(wg *sync.WaitGroup) {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Routine 2 completed")
	wg.Done()
}

func goRoutine3(wg *sync.WaitGroup) {
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Routine 3 completed")
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	wg := &sync.WaitGroup{}
	/*
		wg.Add(1)
		go goRoutine1(wg)
		wg.Add(1)
		go goRoutine2(wg)
		wg.Add(1)
		go goRoutine3(wg)
	*/
	wg.Add(3)
	go goRoutine1(wg)
	go goRoutine2(wg)
	go goRoutine3(wg)
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(start, end, end-start)
	fmt.Println("All the goroutines completed")
}

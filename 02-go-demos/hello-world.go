//package declaration
package main

//import dependency package
import (
	"fmt"
	"time"
)

//package level variables

//package init function

//main function
func main() {
	/*
		var message string
		message = "Hello World!"
	*/

	/*
		var message string = "Hello World!"
		fmt.Println(message)
	*/

	/*
		var name string = "Magesh"
		var greeting string = "Have a nice day!"
	*/

	/*
		var (
			name     string
			greeting string
		)
	*/
	/*
		var (
			name, greeting string
		)
	*/
	/*
		name = "Magesh"
		greeting = "Have a nice day!"
	*/
	/*
		var (
			name     string = "Magesh"
			greeting string = "Have a nice day!"
		)
	*/
	/*
		var (
			name, greeting string = "Magesh", "Have a nice day!"
		)
	*/
	/*
		name, greeting := "Magesh", "Have a nice day!"
	*/

	/*
		const name string = "magesh"
		const greeting string = "Have a nice day!"
	*/
	/*
		const (
			name     string = "Magesh"
		 	greeting string = "Have a nice day!"
		)
	*/
	const pi = 3.14
	const (
		name, greeting string = "Magesh", "Have a nice day!"
	)
	message := greet(name, greeting)

	fmt.Println(message)

	/*
		no := 101
		if no%2 == 0 {
			fmt.Printf("%d is even\n", no)
		} else {
			fmt.Printf("%d is odd\n", no)
		}
		fmt.Println(no)
	*/

	//IF
	if no := 100; no%2 == 0 {
		fmt.Printf("%d is even\n", no)
	} else {
		fmt.Printf("%d is odd\n", no)
	}

	//For
	/*
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
	*/

	// Without Pre/post statements
	/*
		sum := 1
		for ; sum < 1000 ; {
			sum += sum
		}
	*/

	//Using 'for' like a 'while' loop
	/*
		sum := 1
		for sum < 1000 {
			sum += sum
		}
	*/

	//indefinite for
	/*
		for {
			// do something in the loop indefinitely
		}
	*/

	//switch case
	now := time.Now().Unix()
	mins := now % 2
	switch mins {
	case 0:
		fmt.Println("Even")
	/* case 1:
	fmt.Println("Odd") */
	case 3 - 2:
		fmt.Println("Odd")
	}

	score := 5
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	}

	n := 4
	switch n {
	case 0:
		fmt.Println("Zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		fallthrough
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
		fallthrough
	case 9:
		fmt.Println("is <= 9")
		fallthrough
	default:
		fmt.Println("Try again!")
	}

	switch n {
	default:
		fmt.Println("Default")
	}

}

//variadic functions

//other functions
func greet(name string, greeting string) string {
	//fmt.Printf("Hi %s, %s\n", name, greeting)
	return fmt.Sprintf("Hi %s, %s\n", name, greeting)
}

func fn() {
	f2 := func() {

	}
	f2()
}

//go run hello-world.go

//package declaration
package main

//import dependency package
import "fmt"

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
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	//indefinite for
	/*
		for {
			// do something in the loop indefinitely
		}
	*/

}

//other functions
func greet(name string, greeting string) string {
	//fmt.Printf("Hi %s, %s\n", name, greeting)
	return fmt.Sprintf("Hi %s, %s\n", name, greeting)
}

//go run hello-world.go

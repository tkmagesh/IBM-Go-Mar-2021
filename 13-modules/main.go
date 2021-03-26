package main

import (
	"empapp/dept"
	"empapp/emp"
	"fmt"
)

func main() {
	empsChan := make(chan emp.Employees)
	deptChan := make(chan dept.Departments)
	go emp.ParseAsync("emp.dat", empsChan)
	go dept.ParseAsync("dept.dat", deptChan)
	employees := <-empsChan
	departments := <-deptChan
	empWithDeparts, _ := Merge(employees, departments)
	empWithDeparts.Serialize("result.csv")
	fmt.Println("Done")
}

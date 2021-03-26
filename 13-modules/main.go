package main

import (
	"empapp/emp"
	"fmt"
)

func main() {
	employees, err := emp.ParseCSV("emp.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(employees.avgSalary())
}

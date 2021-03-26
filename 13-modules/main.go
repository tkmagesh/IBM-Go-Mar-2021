package main

import (
	"empapp/emp"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	employees, err := emp.ParseCSV("emp.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	color.Green(strconv.FormatFloat(employees.AvgSalary(), 'E', -1, 64))
}

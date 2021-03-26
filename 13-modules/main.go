package main

import (
	"empapp/dept"
	"empapp/emp"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type EmpWithDepart struct {
	Id     int
	Name   string
	Salary float64
	Dept   string
}

func merge(employees emp.Employees, departments dept.Departments) ([]*EmpWithDepart, error) {
	empWithDeparts := make([]*EmpWithDepart, 0)
	for _, emp := range employees {
		empWithDepartment := &EmpWithDepart{
			Id:     emp.Id,
			Name:   emp.Name,
			Salary: emp.Salary,
			Dept:   departments[emp.DeptId],
		}
		empWithDeparts = append(empWithDeparts, empWithDepartment)
	}
	return empWithDeparts, nil
}

func main() {
	empsChan := make(chan emp.Employees)
	deptChan := make(chan dept.Departments)
	go func() {
		employees, err := emp.ParseCSV("emp.dat")
		if err != nil {
			fmt.Println(err)
			return
		}
		empsChan <- employees
	}()
	go func() {
		departments, err := dept.ParseCSV("dept.dat")
		if err != nil {
			fmt.Println(err)
			return
		}
		deptChan <- departments
	}()
	employees := <-empsChan
	departments := <-deptChan
	empWithDeparts, _ := merge(employees, departments)

	file, err := os.Create("result.csv")
	if err != nil {
		fmt.Println("Unable to create Files")
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, empWithDept := range empWithDeparts {
		salValue := strconv.FormatFloat(empWithDept.Salary, 'E', -1, 64)
		value := []string{strconv.Itoa(empWithDept.Id), empWithDept.Name, salValue, empWithDept.Dept}
		err := writer.Write(value)
		if err != nil {
			fmt.Println("Some error during writing")
			return
		}
	}
	fmt.Println("Done")
}

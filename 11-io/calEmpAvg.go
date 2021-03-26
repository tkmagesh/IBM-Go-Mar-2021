package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	Id     int
	Name   string
	Salary float64
	DeptId int
}

type Employees []*Employee

func main() {
	employees, err := parseEmployees("emp.dat")
	fmt.Println(employees)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(employees.avgSalary())

}

func parseEmployees(fileName string) (Employees, error) {
	employees := make([]*Employee, 0)
	inputHandle, inputError := os.Open(fileName)
	defer inputHandle.Close()
	if inputError != nil {
		fmt.Println(inputError)
		return nil, inputError
	}
	lines, err := csv.NewReader(inputHandle).ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		employee, err := parseEmployeeData(line)
		if err == nil {
			employees = append(employees, employee)
		}
	}
	return employees, nil
}
func (employees Employees) totalSalary() float64 {
	total := float64(0)
	for _, employee := range employees {
		total += employee.Salary
	}
	return total
}

func (employees Employees) avgSalary() float64 {
	return employees.totalSalary() / float64(len(employees))
}

func parseEmployeeData(line []string) (*Employee, error) {
	empSalary, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		return nil, err
	}
	empId, err := strconv.Atoi(line[0])
	if err != nil {
		return nil, err
	}
	deptId, err := strconv.Atoi(line[3])
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	emp := &Employee{
		Id:     empId,
		Name:   line[1],
		Salary: empSalary,
		DeptId: deptId,
	}
	return emp, nil
}

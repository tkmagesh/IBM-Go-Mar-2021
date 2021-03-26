package emp

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Employees []*Employee

func ParseAsync(fileName string, ch chan Employees) {
	employees, err := ParseCSV("emp.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- employees
}

func ParseCSV(fileName string) (Employees, error) {
	employees := make([]*Employee, 0)
	inputHandle, inputError := os.Open(fileName)
	defer inputHandle.Close()
	if inputError != nil {
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

func (employees Employees) TotalSalary() float64 {
	total := float64(0)
	for _, employee := range employees {
		total += employee.Salary
	}
	return total
}

func (employees Employees) AvgSalary() float64 {
	return employees.TotalSalary() / float64(len(employees))
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

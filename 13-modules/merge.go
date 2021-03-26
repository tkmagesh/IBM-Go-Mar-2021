package main

import (
	"empapp/dept"
	"empapp/emp"
)

func Merge(employees emp.Employees, departments dept.Departments) (EmpWithDeparts, error) {
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

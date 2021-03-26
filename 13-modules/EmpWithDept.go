package main

import (
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

type EmpWithDeparts []*EmpWithDepart

func (empWithDeparts EmpWithDeparts) Serialize(fileName string) {
	file, err := os.Create(fileName)
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
}

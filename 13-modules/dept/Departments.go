package dept

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Departments map[int]string

func ParseCSV(fileName string) (Departments, error) {
	departments := Departments{}
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
		value, _ := strconv.Atoi(line[0])
		departments[value] = line[1]
	}
	return departments, nil
}

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContents, err := ioutil.ReadFile("sample.dat")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fileContents))
	err = ioutil.WriteFile("sample-copy.dat", fileContents, 0x777)
	if err != nil {
		fmt.Println(err)
	}

}

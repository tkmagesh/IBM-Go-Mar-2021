package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputHandle, inputError := os.Open("sample.dat")
	/* defer func() {
		fmt.Printf("The first defer")
		inputHandle.Close()
	}() */
	defer inputHandle.Close()
	defer func() {
		fmt.Printf("The second defer")
	}()
	if inputError != nil {
		fmt.Println(inputError)
		return
	}
	inputReader := bufio.NewReader(inputHandle)
	for {
		line, err := inputReader.ReadString('\n')
		fmt.Println("Line :")
		fmt.Println(line)
		if err == io.EOF {
			break
		}

	}

}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func getResponse(ch chan *http.Response) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- resp
}

func main() {
	respCh := make(chan *http.Response)
	go getResponse(respCh)
	timeout := time.After(100 * time.Millisecond)
	for {
		select {
		case res := <-respCh:
			io.Copy(os.Stdout, res.Body)
			return
		case <-timeout:
			fmt.Println("timeout!!")
			return
		}
	}
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

const HOST_NAME = "http://62.84.125.40"

func DoGet(url string) (int, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode, body
}

func queryMock() {
	const URL = HOST_NAME + "/comments?postId=1"
	status, body := DoGet(URL)
	fmt.Printf("status: %d, body: %s\n", status, body)
}

func main() {
	queryMock()
}

// We want to proxy to https://jsonplaceholder.typicode.com/comments?postId=1

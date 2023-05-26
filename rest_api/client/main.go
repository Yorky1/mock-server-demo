package main

import (
	"fmt"
	"io"
	"net/http"
)

const MOCK_SERVER_HOST = "http://62.84.125.40"
const EXTERNAL_HOST = "http://51.250.91.250:8080"

func queryMock() {
	const URL = MOCK_SERVER_HOST + "/data"
	status, body := DoGet(URL)
	fmt.Printf("status: %d, body: %s\n", status, body)
}

func main() {
	queryMock()
}

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

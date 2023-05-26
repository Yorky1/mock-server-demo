package main

import (
	"fmt"
	"log"
	"net/http"
)

func data(w http.ResponseWriter, req *http.Request) {
	// Some hard work

	sum := 0
	for i := 0; i <= 1000000000; i += 1 {
		sum += i
	}
	fmt.Fprintf(w, fmt.Sprintf("production data: %d", sum))
}

func main() {
	http.HandleFunc("/data", data)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

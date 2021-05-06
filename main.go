package main

import (
	"fmt"
	"net/http"
)

func hello_page(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello back!")
	fmt.Fprintf(w, "hello front!")
}

func handleRequest() {
	http.HandleFunc("/", hello_page)
	http.HandleFunc("/foo/", hello_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}

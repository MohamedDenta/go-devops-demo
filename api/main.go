package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, DevOps with Go!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Println("Server started on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

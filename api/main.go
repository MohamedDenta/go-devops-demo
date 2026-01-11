package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, DevOps with Go!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Println("Server started on :9090")
	log.Println("ENV PORT is set to:", getenv("PORT", "not set"))
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func getenv(s1, s2 string) string {
	s := os.Getenv(s1)
	if s == "" {
		return s2
	}
	return s
}

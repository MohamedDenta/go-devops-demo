package main

import (
	"log"

)

func main() {
	log.Println("Server started on :9090")
	log.Println("ENV PORT is set to:", getenv("PORT", "not set"))

	startServer()
}



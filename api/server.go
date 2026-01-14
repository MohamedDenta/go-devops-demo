package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hello, DevOps with Go!") }
func startServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	registerHealthHandlers(mux)

	srv := &http.Server{
		Addr:    ":" + getenv("PORT", "9090"),
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown (will complete in next task)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

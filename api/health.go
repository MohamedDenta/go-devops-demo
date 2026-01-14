package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

var ready atomic.Bool

func registerHealthHandlers(mux *http.ServeMux) {
	ready.Store(true)

	mux.HandleFunc("/health/live", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Liveness check")
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		if !ready.Load() {
			http.Error(w, "not ready", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/health/startup", func(w http.ResponseWriter, r *http.Request) {
		if !ready.Load() {
			http.Error(w, "starting", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

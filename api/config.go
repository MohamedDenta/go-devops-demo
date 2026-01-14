package main

import (
	"os"
)

func getenv(s1, s2 string) string {
	s := os.Getenv(s1)
	if s == "" {
		return s2
	}
	return s
}

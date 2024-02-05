package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Failed to start HTTP server.", err)
		os.Exit(1)
	}
}

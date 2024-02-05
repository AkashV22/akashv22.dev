package main

import (
	"bytes"
	"log/slog"
	"net/http"
	"os"
	"text/template"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/pages/home.html"))

		var buf bytes.Buffer

		if err := tmpl.Execute(&buf, nil); err != nil {
			slog.Error("Error rendering HTML template.", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Error starting HTTP server.", err)
		os.Exit(1)
	}
}

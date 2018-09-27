package main

import (
	"fmt"
	"net/http"
)

const version = "1.0.1"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Version: %s", version)
	})

	http.ListenAndServe("localhost:8080", nil)
}

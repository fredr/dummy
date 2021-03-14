package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":3000", nil)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", any)
	http.ListenAndServe(":3000", nil)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func any(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	w.Write([]byte("OK"))
}

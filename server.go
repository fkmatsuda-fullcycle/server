package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle!!!</h1>"))
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", hello)

	fmt.Println("Server started...")

	go func() {
		for true {
			fmt.Println("tic...")
			time.Sleep(time.Second)
			fmt.Println("tac...")
			time.Sleep(time.Second)
		}
	}()

	http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello %s!!! You are %s years old.", name, age)

	fmt.Println("Hello World!")

}

func healthz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}

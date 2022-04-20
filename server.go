package main

import (
	"fmt"
	"math/rand"
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

	s := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(s)
	num := rnd.Intn(9999)

	fmt.Fprintf(w, "<h1>Hello %s!!! You are %s years old.</h1><h3>My number: %d</h3>", name, age, num)

	fmt.Println("Hello World!", num)

}

func healthz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}

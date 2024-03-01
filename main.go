package main

import (
	"fmt"
	"log"
	"net/http"
)

var port string = ":4000"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Golang"))
}

func main() {
	fmt.Println("Hello Snippet")

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting server on " + port)
	//ListenAndServe takes the port and the mux
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

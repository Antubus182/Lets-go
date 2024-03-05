package main

import (
	"fmt"
	"log"
	"net/http"
)

var port string = ":4000"

func main() {
	fmt.Println("Hello Snippet")

	//Using a serveMux is good practise because we can define all routes here instead of having many http handlefuncs
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on " + port)
	//ListenAndServe takes the port and the mux
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

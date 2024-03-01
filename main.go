package main

import (
	"fmt"
	"log"
	"net/http"
)

var port string = ":4000"

func home(w http.ResponseWriter, r *http.Request) {

	// because "/" is a subtree or catchall, we need to manually restrict access
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Print("Accessed illegal, returned not found")
		return
	}
	w.Write([]byte("Hello from Golang"))
	log.Print("accessed home")
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Write a snippet"))
	log.Print("accessed view")
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new Snippet"))
	log.Print("accessed create")
}

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

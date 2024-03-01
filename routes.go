package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	// because "/" is a subtree or catchall, we need to manually restrict access
	if r.URL.Path != "/" {
		http.Error(w, "Unregistered path", http.StatusTeapot) //Go supports 418 I'm a teapot
		log.Print("Accessed illegal, returned not found")
		return
	}
	w.Write([]byte("Hello from Golang"))
	log.Print("accessed home")
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		//alternative to writeheader and then write is using Error
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Write a snippet"))
	log.Print("accessed view")
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	//Restricting this to only POST method
	if r.Method != "POST" {
		w.WriteHeader(405) //405= Method Not Allowed
		//if we want a non 200 header we need to call this before the w.Write call
		w.Write([]byte("Method not Allowed"))
		log.Print("non Post request on snippetCreate")
		return
	}

	w.Write([]byte("Create a new Snippet"))
	log.Print("accessed create")
}

func JsonReturn(w http.ResponseWriter, r *http.Request) {
	//if the response is json we need to manually set the content type in the header to json
	w.Header().Set("Content-Type", "application/json")
	//Header().Set overrides all and can thus be used once, for multiple settins use add
	w.Header().Add("Cache-Control", "max-age=31536000")
	w.Write([]byte(`{"name":"Antubus"}`))
}

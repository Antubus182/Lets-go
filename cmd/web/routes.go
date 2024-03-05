package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	// because "/" is a subtree or catchall, we need to manually restrict access
	if r.URL.Path != "/" {
		http.Error(w, "Unregistered path", http.StatusTeapot) //Go supports 418 I'm a teapot
		log.Print("Accessed illegal, returned not found")
		return
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Parsing error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		//alternative to writeheader and then write is using Error
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Showing a snippet wit ID %d...", id)
	log.Print("accessed view with ID")
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

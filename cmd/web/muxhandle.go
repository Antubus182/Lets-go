package main

import "net/http"

func (app *application) muxroutes() *http.ServeMux {
	fileServer := http.FileServer(http.Dir("./ui/Static/"))

	//Using a serveMux is good practise because we can define all routes here instead of having many http handlefuncs
	mux := http.NewServeMux()
	mux.Handle("/Static/", http.StripPrefix("/Static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux

}

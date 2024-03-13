package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"npi/snippetbox/internal/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// because "/" is a subtree or catchall, we need to manually restrict access
	if r.URL.Path != "/" {
		app.notFound(w)
		app.logger.Warn("Accessed illegal, returned not found", slog.String("Path", r.URL.Path))
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}
	/*
		templates := []string{
			"./ui/html/base.tmpl",
			"./ui/html/pages/home.tmpl",
			"./ui/html/partials/nav.tmpl",
		}

		ts, err := template.ParseFiles(templates...)
		if err != nil {
			app.logger.Error(err.Error())
			app.serverError(w, r, err)
			return
		}
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Print(err.Error())
			app.serverError(w, r, err)
		}
	*/
	app.logger.Debug("Accessed Home page")
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		//alternative to writeheader and then write is using Error
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	app.logger.Debug("accessed view with ID")
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	//Restricting this to only POST method
	if r.Method != "POST" {
		w.WriteHeader(405) //405= Method Not Allowed
		//if we want a non 200 header we need to call this before the w.Write call
		w.Write([]byte("Method not Allowed"))
		app.logger.Warn("non Post request on snippetCreate")
		return
	}

	app.logger.Debug("accessed create")
	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}

func JsonReturn(w http.ResponseWriter, r *http.Request) {
	//if the response is json we need to manually set the content type in the header to json
	w.Header().Set("Content-Type", "application/json")
	//Header().Set overrides all and can thus be used once, for multiple settins use add
	w.Header().Add("Cache-Control", "max-age=31536000")
	w.Write([]byte(`{"name":"Antubus"}`))
}

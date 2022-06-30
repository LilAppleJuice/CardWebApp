package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func route() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()

	//var usrC string

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/homepage", homepageHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "homepage.html", nil)
}

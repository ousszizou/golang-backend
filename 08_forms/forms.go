package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/ousszizou/go_backend/08_forms/message"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// SERVER PORT
const (
	PORT = ":8080"
)

func render(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}

	t.Execute(w, data)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		render(w, "./templates/contact.html", nil)
	}
	if r.Method == "POST" {
		msg := &message.Message{
			Email:   r.PostFormValue("email"),
			Content: r.PostFormValue("content"),
		}

		if msg.Validate() == false {
			render(w, "./templates/contact.html", msg)
			return
		}

		http.Redirect(w, r, "/confirm", http.StatusSeeOther)
	}
}

func confirmHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "./templates/confirm.html", nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/contact", contactHandler).Methods("GET", "POST")
	r.HandleFunc("/confirm", confirmHandler)

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(PORT, loggedRouter)
}

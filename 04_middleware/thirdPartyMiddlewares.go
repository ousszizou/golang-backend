package main

import (
	"fmt"
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index page")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/home", homeHandler)

	// http.Handle("/", r)

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)

	http.ListenAndServe(":8000", loggedRouter)
}

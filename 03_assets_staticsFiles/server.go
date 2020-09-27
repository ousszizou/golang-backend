package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}

	t.Execute(w, nil)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/home", homeHandler)

	// Explanation:

	// FileServer() is told the root of static files is "assets".
	// We want the URL to start with "/assets/".

	// So if someone requests "/assets/scripts/main.js", we want the server to send the file "/assets/scripts/main.js".

	// So we have to strip "/assets/" from the URL, and the remaining will be the relative path compared to the root folder "assets".

	// assets/scripts/main.js (with StripPrefix)
	// assets/assets/scripts/main.js (without StripPrefix)

	router.
		PathPrefix("/assets/").
		Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8000", router)
}

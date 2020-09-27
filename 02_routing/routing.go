package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]   // the book's title
	number := vars["number"] // the page's number

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, number)
}

func main() {

	// Creating a new Router
	r := mux.NewRouter()

	// :8080/books/learn_go/page/200
	// :8080/books/learn_java/page/301
	// :8080/books?title=learn_go&page=200
	r.HandleFunc("/books/{title}/page/{number}", handler)

	http.ListenAndServe(":8080", r)
}

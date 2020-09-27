package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	// "text/template"
)

// Todo type
type Todo struct {
	Title string
	Done bool
}

// TodoPage type
type TodoPage struct {
	PageTitle string
	Todos []Todo
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To Index Page")	
}

func todosHandler(w http.ResponseWriter, r *http.Request) {

	data := TodoPage{
		PageTitle: "All Todos",
		Todos: []Todo{
			{Title: "Learn Go", Done: false},
			{Title: "Learn Python", Done: true},
			{Title: "Learn JavaScript", Done: true},
		},
	}

	t, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}

	t.Execute(w,data)
}

func main()  {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/todos", todosHandler)
	http.ListenAndServe(":8080", nil)
}

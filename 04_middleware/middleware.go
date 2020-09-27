// // When you're building a web application there's probably some shared functionality that you want to run for many (or even all) HTTP requests.

// // One way of organising this shared functionality is to set it up as middleware – self-contained code which independently acts on a request before or after your normal application handlers.

// // Router => Middleware Handler => Application Handler

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func now() string {
// 	return time.Now().Format(time.Stamp) + " "
// }

// func handler1(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("handler 1")
// }

// func handler2(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("handler 2")
// }

// func logger(f http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(now() + "Before")
// 		defer fmt.Println(now() + "after")

// 		f(w, r)
// 	}
// }

// func handler3(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("handler 3")
// }


// func main() {
// 	http.HandleFunc("/h1", logger(handler1))
// 	http.HandleFunc("/h2", logger(handler2))
// 	http.HandleFunc("/h3", logger(handler3))

// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"fmt"
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello every body</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1> FAQ PAGE </h1>
	
	`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HomeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "<h1>PAGE "+r.URL.Path+" NOT FOUND</h1>", http.StatusNotFound)
	}
}
func main() {
	var router Router
	fmt.Println("Starting Server on port 8080")
	http.ListenAndServe(":8080", router)
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var empty interface{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, empty)
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("serving on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

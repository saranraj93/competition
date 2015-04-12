package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var listen = ":8080"

var templates = template.Must(template.ParseFiles("templates/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	fmt.Printf("")

	log.Println("Listening on", listen)
	http.ListenAndServe(listen, nil)
}

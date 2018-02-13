package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("").ParseFiles("templates/index.html", "templates/layout.html")
	greeting := os.Getenv("GREETING")
	err = tpl.ExecuteTemplate(w, "layout", greeting)
	if err != nil {
		log.Printf("Error serving Admin %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

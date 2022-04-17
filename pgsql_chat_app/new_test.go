package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/gin-gonic/gin"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("login.html"))
	// t, _ := template.ParseFiles("login.html")
	tmpl.ExecuteTemplate(w, "Index", nil)
}

func main() {
	log.Println("Server started on: http://localhost:9090")
	http.HandleFunc("/", indexHandler)

	// http.HandleFunc("/show", Show)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/update", Update)
	// http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":9090", nil)
}

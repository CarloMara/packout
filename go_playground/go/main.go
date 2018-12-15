package main

import(
	// "io/ioutil"
	"net/http"
	"html/template"
	"log"
	// "regexp"
	// "errors"
	// "os"
)

var templates = template.Must(template.ParseFiles("static/templates/root.html"))

type Global_data struct{
	Title string
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "root.html", &Global_data{Title:"TEST TEMPLATE PARSING"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func cssHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "static/css/wing.css")
}


func main() {
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/static/css/wing.css", cssHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}
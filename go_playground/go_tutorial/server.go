package main

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"log"
	"regexp"
	"errors"
	"os"
)

var templates = template.Must(template.ParseFiles("static/tmpl/edit.html", "static/tmpl/view.html", "static/tmpl/root.html"))
var validPath = regexp.MustCompile("/([a-zA-Z0-9]+)$")

// page struct
type Page struct {
	Title string
	Body []byte
}

var logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

//function on page that saves the content on a file
//could be replaced with a database connection
func (p *Page)save() error {
	filename := "data/" + p.Title + ".txt"
	err := ioutil.WriteFile(filename, p.Body, 0600)
	if err != nil {
		return errors.New("could not create file. abort")
	}
	return nil
}

func loadPage(title string) (*Page, error){
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTitle( w http.ResponseWriter, r *http.Request) (string, error){
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Title Page")
	}
	return m[2], nil
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string){
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/tmpl/root.html")
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		logger.Printf("request url path: %s makeHandler subpath %s", r.URL.Path, m)

		if m == nil{
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[1])
	}
}




func main() {
	http.HandleFunc("/", rootHandler)
	
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
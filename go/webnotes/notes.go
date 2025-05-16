package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/pypaut/slices"
)

var templateDir = "templates"
var validPath = regexp.MustCompile("^/(edit|save|view|new)/([a-zA-Z0-9_]+)$")

var templates = template.Must(template.ParseFiles(
	templateDir+"/root.html",
	templateDir+"/view.html",
	templateDir+"/edit.html",
	templateDir+"/new.html",
))

type Note struct {
	Title string
	Body  []byte
}

func (n *Note) save() error {
	filename := "notes/" + n.Title + ".txt"
	return os.WriteFile(filename, n.Body, 0600)
}

func loadNote(title string) (*Note, error) {
	filename := "notes/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Note{Title: title, Body: body}, nil
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}

	return m[2], nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, n any) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", n); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	noteDirs, err := os.ReadDir("./notes")
	if err != nil {
		log.Fatal(err)
	}

	// Convert to slice of Note
	notes, err := slices.Map(noteDirs, func(d os.DirEntry) (*Note, error) {
		noteTitle := strings.TrimSuffix(d.Name(), ".txt")
		n, err := loadNote(noteTitle)
		if err != nil {
			return nil, err
		}

		return n, nil
	})

	if err != nil {
		log.Fatal(err)
	}

	renderTemplate(w, "root", notes)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	n, err := loadNote(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", n)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	n, err := loadNote(title)
	if err != nil {
		n = &Note{Title: title}
	}

	renderTemplate(w, "edit", n)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "new.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	body := r.FormValue("body")
	n := &Note{Title: title, Body: []byte(body)}

	err = n.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	n := &Note{Title: title, Body: []byte(body)}

	err := n.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/save/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/new/", newHandler)
	http.HandleFunc("/create/", createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

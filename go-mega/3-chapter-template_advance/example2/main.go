package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// User struct
type User struct {
	Username string
}

// Post
type Post struct {
	User
	Body string
}

// IndexViewModel
type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

// PopulateTemplates func
// Create map template name to template.Template
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{Username: "zhengqilin"}
		u2 := User{Username: "metrosir"}

		posts := []Post{
			Post{User: u1, Body: "good"},
			Post{User: u2, Body: "good"},
		}
		v := IndexViewModel{Title: "HomePage", User: u1, Posts: posts}
		templates := PopulateTemplates()
		templates["index.html"].Execute(w, &v)
	})
	http.ListenAndServe(":1021", nil)
}

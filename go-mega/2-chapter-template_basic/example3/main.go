package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

type Post struct {
	User User
	Body string
}

type IndexViewMode struct {
	Title string
	User  User
	Posts []Post
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{Username: "metrosir"}
		u2 := User{Username: "zhengqilin"}

		posts := []Post{
			Post{User: u1, Body: "good"},
			Post{User: u2, Body: "good"},
		}

		v := IndexViewMode{Title: "HomePage", User: u1, Posts: posts}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &v)
	})
	http.ListenAndServe(":1201", nil)
}

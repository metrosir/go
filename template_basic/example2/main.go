package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

type IndexViewModel struct {
	Title string
	User  User
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "metrosir"}
		v := IndexViewModel{Title: "HomePage", User: user}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &v)
	})
	http.ListenAndServe(":1201", nil)
}

package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "zheng"}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &user)
	})
	http.ListenAndServe(":1201", nil)
}

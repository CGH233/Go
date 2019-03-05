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
	User User
}

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{"CGH"}
		v := IndexViewModel{"Home",user}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &v)

	})
	http.ListenAndServe(":2000", nil)
}
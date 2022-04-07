package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "any code you want in this func")
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// data := struct {
	// 	Method      string
	// 	URL         *url.URL
	// 	Submissions map[string][]string
	// }{
	// 	r.Method,
	// 	r.URL,
	// 	r.Form,
	// }

	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	fmt.Println("sever runing on http://localhost:8080")
	http.ListenAndServe(":8080", d)
}

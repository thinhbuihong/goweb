package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "any code you want in this func")
	io.WriteString(w, "doggy doggy doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "any code you want in this func")
	io.WriteString(w, "cat cat cat")
}

func f(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "fish fish fish")
}

func main() {
	var d hotdog
	var c hotcat
	// mux := http.NewServeMux()
	// mux.Handle("/dog/", d)
	// mux.Handle("/cat/", c)

	http.Handle("/dog/", d)
	http.Handle("/cat/", c)

	http.HandleFunc("/fish/", f)

	fmt.Println("sever runing on http://localhost:8080")
	http.ListenAndServe(":8080", nil) //default serve mux
}

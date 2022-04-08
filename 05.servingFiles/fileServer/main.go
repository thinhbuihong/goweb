package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.Handle("/resources/", http.StripPrefix("/resources",
		http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":3000", nil))

	http.Handle("/", http.NotFoundHandler())
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/resources/toby.jpg"/>`)
	// io.WriteString(w, `<img src="/toby.jpg"/>`)
}

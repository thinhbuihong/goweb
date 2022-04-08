package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogpic)
	http.ListenAndServe(":3000", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/toby.jpg"/>`)
}

func dogpic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 400)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

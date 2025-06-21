package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})
	http.HandleFunc("/js/actors.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "site/js/actors.js")
	})
	http.HandleFunc("/css/STYLES.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "site/css/STYLES.css")
	})
	http.HandleFunc("/media/nofoto.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "site/media/nofoto.jpg")
	})
	http.HandleFunc("/actors", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "site/actors.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Index Page")
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}

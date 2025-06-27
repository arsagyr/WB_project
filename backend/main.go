package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"backend/internal/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func runServer() {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=testactors sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	handlers.DB = db
	defer handlers.DB.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/create", handlers.CreateHandler)
	router.HandleFunc("/edit/{id:[0-9]+}", handlers.EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", handlers.EditHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", handlers.DeleteHandler)

	http.Handle("/", router)
	fmt.Println("Server is listening...")

	http.ListenAndServe("localhost:8181", nil)
}

func main() {
	runServer()

}

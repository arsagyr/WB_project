package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/arry/WB_project/internal/handlers"
	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=postgres password=password dbname=actorsdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	http.HandleFunc("/css/STYLES.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "site/css/STYLES.css")
	})
	http.HandleFunc("/media/nofoto.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "site/media/nofoto.jpg")
	})

	http.HandleFunc("/create", handlers.CreateHandler2)
	http.HandleFunc("/", handlers.TableHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)

}

// func main() {

// 	d := displayed.Displayed{
// 		Familyname: "Bor",
// 		Givenname:  "Alex",
// 		Nation:     " - ",
// 		Honorar:    1,
// 		Number:     1,
// 	}
// 	displayed.Insert(d)
// 	displayed.PrintDisplayed()
// }
